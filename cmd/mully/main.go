package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
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

	tmpLog := initConsoleLogger(nil)

	conf, err := getConfig()
	if err != nil {
		tmpLog.Logger.Fatal().Err(err).Msg("missing configuration")
	}

	checkProto(conf.logLevel == zerolog.TraceLevel)
	client, err := NewRPCClient(conf)
	if err != nil {
		tmpLog.Logger.Fatal().Err(err).Msg("failed to connect to the mullvad daemon")
	}

	defer func() {
		client.log.Logger.Warn().Msg("closing connection to the mullvad daemon (main defer)")
		if err = client.Close(); err != nil {
			client.log.Logger.Error().Err(err).Msg("failed to close connection to the mullvad daemon")
			return
		}
		client.log.Logger.Debug().Msg("connection closed")
	}()

	client.log.Logger.Debug().Msg("client initialized")

	client.log.Logger.Debug().Msg("connecting...")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(conf.timeout))
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		client.log.Logger.Fatal().Err(err).Msg("failed to connect to the mullvad daemon")
	}

	<-sigChan
}
