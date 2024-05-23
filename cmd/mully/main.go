package main

import (
	"context"
	"encoding/hex"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"

	"git.tcp.direct/kayos/mully/pkg/mullvad_mgmt"
)

func printAccountInfo(client *RPCClient) {
	/*var (
		accountID     string
		accountExpiry time.Time
		err           error
	)

	accountID, accountExpiry, err = client.GetAccountInfo(ctx)
	if err != nil {
		client.log.Logger.Fatal().Err(err).Msg("failed to get account info")
	}
	client.log.Logger.Trace().Msg("got account info")
	client.log.Logger.Info().Msgf("account id: %s, expiry: %s", accountID, accountExpiry.String())*/
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	conf, err := getConfig()
	if err != nil {
		println(err.Error())
		return
	}

	checkProto(conf.logLevel == zerolog.TraceLevel)
	client, err := NewRPCClient(conf)
	if err != nil {
		println(err.Error())
		return
	}

	zlog := client.log.Logger

	defer func() {
		zlog.Warn().Msg("signal caught, shutting down")
		if err = client.Close(); err != nil {
			zlog.Error().Err(err).Msg("failed to close connection to the mullvad daemon")
			return
		}
		zlog.Debug().Msg("connection closed")
	}()

	zlog.Trace().Msg("client initialized")

	zlog.Debug().Msg("connecting...")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(conf.timeout))
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		zlog.Fatal().Err(err).Msg("failed to connect to the mullvad daemon")
	}

	deviceState, err := client.mgmtClient.GetDevice(ctx, &emptypb.Empty{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to get device state")
	}

	var accountAndDevice *mullvad_mgmt.AccountAndDevice

	if accountAndDevice = deviceState.GetDevice(); accountAndDevice == nil {
		zlog.Fatal().Msg("device state is nil")
		return // unreachable, but the nilness analyzer (linter) was confused.
	}

	if client.config.Account = accountAndDevice.GetAccountToken(); client.config.Account == "" {
		zlog.Fatal().Msg("account token is empty")
	}

	zlog.Info().Msgf("account token: %s", client.config.Account)

	if client.device = accountAndDevice.GetDevice(); client.device == nil {
		zlog.Fatal().Msg("device is empty")
	}

	zlog.Info().Str("ID", client.device.GetId()).Send()
	zlog.Info().Str("name", client.device.GetName()).Send()
	zlog.Info().Str("pubkey", hex.EncodeToString(client.device.GetPubkey())).Send()
	zlog.Info().Bool("hijack_dns", client.device.GetHijackDns()).Send()
	zlog.Info().Time("created", client.device.GetCreated().AsTime()).Send()

	<-sigChan
}
