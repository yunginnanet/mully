package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"git.tcp.direct/kayos/zwrap"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"git.tcp.direct/kayos/mully/pkg/mullvad_mgmt"
)

func initConsoleLogger() *zwrap.Logger {
	zl := zerolog.New(zerolog.ConsoleWriter{
		Out:         os.Stderr,
		FormatLevel: zwrap.LogLevelFmt(runtime.GOOS == "windows"),
	}).With().Timestamp().Logger()
	return zwrap.Wrap(zl)
}

type RPCClient struct {
	*grpc.ClientConn
	log    *zwrap.Logger
	hooker zerolog.Hook
	config *Config
}

// because there's no need to export `Run`, this is a hack
type logHooker struct {
	rpc *RPCClient
}

func (h *logHooker) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	h.rpc.hookRun(e, level, msg)
}

func (r *RPCClient) hookRun(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("state", r.GetState().String())
	e.Str("caller", r.CanonicalTarget())
}

func NewRPCClient(c *Config) (*RPCClient, error) {
	client, err := grpc.NewClient(
		"unix:///var/run/mullvad-vpn",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the mullvad daemon: %w", err)
	}
	rpcc := &RPCClient{
		ClientConn: client,
	}
	rpcc.hooker = &logHooker{rpc: rpcc}
	rpcc.log = initConsoleLogger().WithForceLevel(zerolog.DebugLevel)
	rpcc.log.Hook(rpcc.hooker)
	grpclog.SetLoggerV2(rpcc.log)
	return rpcc, nil
}

func (r *RPCClient) Close() error {
	return r.ClientConn.Close()
}

func (r *RPCClient) stateWatcher(ctx context.Context, states chan<- connectivity.State, readyConn chan<- struct{}) {
	log := r.log
	last := connectivity.Idle
	for {
		if ok := r.ClientConn.WaitForStateChange(ctx, last); !ok {
			// expired context
			return
		}

		last = r.GetState()
		states <- last

		switch last {
		case connectivity.Ready:
			readyConn <- struct{}{}
			continue
		case connectivity.TransientFailure:
			log.Logger.Error().Msg("connection in transient failure state")
		case connectivity.Shutdown:
			log.Logger.Warn().Msg("connection in shutdown state")
		default:
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (r *RPCClient) connect(ctx context.Context) error {
	log := r.log

	states := make(chan connectivity.State, 1)
	readyConn := make(chan struct{}, 1)

	go r.stateWatcher(ctx, states, readyConn)

waitLoop:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-states:
			log.Logger.Debug().Msg("state changed")
		case <-readyConn:
			log.Logger.Info().Msg("connection ready")
			break waitLoop
		}
	}

	return nil
}

func (r *RPCClient) Connect(ctx context.Context) (mullvad_mgmt.ManagementServiceClient, error) {
	log := r.log

	if err := r.connect(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to the gRPC endpoint: %w", err)
	}

	serviceClient := mullvad_mgmt.NewManagementServiceClient(r.ClientConn)
	verStrPB, err := serviceClient.GetCurrentVersion(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get current daemon version: %w", err)
	}

	log.Logger.Info().Msgf("current daemon version: %s", verStrPB.Value)
	verPB, err := serviceClient.GetVersionInfo(ctx, &emptypb.Empty{})
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("failed to get daemon version info")
	}
	log.Logger.Info().Msgf("daemon version info: %s", verPB.String())

	return serviceClient, nil
}

type Config struct {
	Account string
}

func getConfig() (*Config, error) {
	if len(os.Args) < 2 && os.Getenv("MULLVAD_ACCOUNT") == "" {
		return nil, fmt.Errorf("no command specified")
	}
	if os.Getenv("MULLVAD_ACCOUNT") != "" {
		return &Config{Account: os.Getenv("MULLVAD_ACCOUNT")}, nil
	}
	return &Config{Account: os.Args[1]}, nil
}

func main() {
	checkProto(true)

	tmpLog := initConsoleLogger()

	conf, err := getConfig()
	if err != nil {
		tmpLog.Logger.Fatal().Err(err).Msg("missing mullvad account key")
	}
	client, err := NewRPCClient(conf)

	if err != nil {
		tmpLog.Logger.Fatal().Err(err).Msg("failed to connect to the mullvad daemon")
	}

	log := client.log
	defer func() {
		if err = client.Close(); err != nil {
			log.Logger.Error().Err(err).Msg("failed to close connection to the mullvad daemon")
			return
		}
		log.Logger.Debug().Msg("connection closed")
	}()

	log.Logger.Debug().Msg("client initialized")

	log.Logger.Debug().Msg("connecting...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var serviceClient mullvad_mgmt.ManagementServiceClient

	if serviceClient, err = client.Connect(ctx); err != nil {
		log.Logger.Fatal().Err(err).Msg("failed to connect to the mullvad daemon")
	}

	var acct *mullvad_mgmt.AccountData

	if acct, err = serviceClient.GetAccountData(ctx, &wrapperspb.StringValue{Value: ""}); err != nil {
		log.Logger.Fatal().Err(err).Msg("failed to get account data")
	}

}
