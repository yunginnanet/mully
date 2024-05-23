package main

import (
	"context"
	"fmt"
	"time"

	"git.tcp.direct/kayos/zwrap"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"

	"git.tcp.direct/kayos/mully/pkg/mullvad_mgmt"
)

type RPCClient struct {
	*grpc.ClientConn
	mgmtClient mullvad_mgmt.ManagementServiceClient
	log        *zwrap.Logger
	hooker     *logHooker
	config     *Config
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

	rpcc.config = c
	rpcc.log = initConsoleLogger(rpcc.hooker, c)
	rpcc.log.Logger.Trace().Msg("setting gRPC logger...")

	grpclog.SetLoggerV2(initConsoleLogger(nil).WithForceLevel(zerolog.TraceLevel))

	rpcc.log.Logger.Trace().Msg("gRPC logger set")
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
	states := make(chan connectivity.State, 1)
	readyConn := make(chan struct{}, 1)

	go r.stateWatcher(ctx, states, readyConn)
	r.ClientConn.Connect()

waitLoop:
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case newState := <-states:
			r.log.Logger.Debug().Str("state", newState.String()).Msg("state changed")
		case <-readyConn:
			r.log.Logger.Info().Msg("connected")
			break waitLoop
		}
	}

	return nil
}

func (r *RPCClient) Connect(ctx context.Context) error {
	zlog := r.log.Logger

	if err := r.connect(ctx); err != nil {
		return fmt.Errorf("failed to connect to the gRPC endpoint: %w", err)
	}

	r.mgmtClient = mullvad_mgmt.NewManagementServiceClient(r.ClientConn)
	verStrPB, err := r.mgmtClient.GetCurrentVersion(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("failed to get current daemon version: %w", err)
	}

	zlog.Trace().Msgf("current daemon version: %s", verStrPB.Value)
	verPB, err := r.mgmtClient.GetVersionInfo(ctx, &emptypb.Empty{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to get daemon version info")
	}
	zlog.Info().Msgf("daemon version info: %s", verPB.String())

	return nil
}

/*func (r *RPCClient) GetAccountInfo(ctx context.Context) (ID string, expiry time.Time, err error) {
	r.log.Trace("getting account info")
	var acct *mullvad_mgmt.AccountData

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if acct, err = r.mgmtClient.
		GetAccountData(ctx, wrapperspb.String(r.config.Account)); err != nil || acct == nil {
		r.log.Logger.Trace().Err(err).Msg("failed to get account data")
		return "", time.Time{}, fmt.Errorf("failed to get account data: %w", err)
	}

	r.log.Logger.Trace().Msgf("account data: %s", acct.String())

	accountID := acct.GetId()
	if accountID == "" {
		return "", time.Time{}, fmt.Errorf("account id is empty")
	}
	accountExpiry := acct.GetExpiry()
	if accountExpiry.AsTime().IsZero() {
		return "", time.Time{}, fmt.Errorf("account expiry is empty")
	}
	if accountExpiry.AsTime().Before(time.Now()) {
		return accountID, accountExpiry.AsTime(), fmt.Errorf("account is expired")
	}

	return accountID, accountExpiry.AsTime(), nil
}
*/
