package main

import (
	"context"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/term"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"git.tcp.direct/kayos/mully/pkg/mullvad_mgmt"
)

func rl(ctx context.Context) (chan string, error) {
	fd := int(os.Stdin.Fd())
	w, h, e := term.GetSize(fd)
	if e != nil {
		return nil, e
	}
	winch := make(chan os.Signal, 1)
	signal.Notify(winch, syscall.SIGWINCH)
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, err
	}
	t := term.NewTerminal(os.Stdin, "mully> ")
	if err = t.SetSize(w, h); err != nil {
		return nil, err
	}

	lines := make(chan string)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line, lineErr := t.ReadLine()
				if lineErr != nil && !errors.Is(lineErr, io.EOF) {
					panic(lineErr)
				}
				lines <- line
			}
		}
	}()

	go func() {
		defer func() {
			if restoreErr := term.Restore(fd, oldState); restoreErr != nil {
				panic(restoreErr)
			}
		}()
		for {
			select {
			case <-winch:
				w, h, e = term.GetSize(fd)
				if e != nil {
					panic(e)
				}
				if err = t.SetSize(w, h); err != nil {
					panic(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return lines, nil
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

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	events, err := client.mgmtClient.EventsListen(ctx, &emptypb.Empty{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to listen for daemon events")
	}

	go func() {
		for {
			de := &mullvad_mgmt.DaemonEvent{}
			// when ctx is canceled, this will return an error, so we don't need to poll ctx.Done()
			if err = events.RecvMsg(de); err != nil {
				if !errors.Is(err, io.EOF) {
					zlog.Error().Err(err).Msg("failed to receive daemon event")
				}
				return
			}
			zlog.Info().Msgf("daemon event: [%T] %v", de, de)
		}
	}()

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

	excludedProcesses, err := client.mgmtClient.GetExcludedProcesses(ctx, &emptypb.Empty{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to get excluded processes")
	}

	tunnelState, err := client.mgmtClient.GetTunnelState(ctx, &emptypb.Empty{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to get tunnel state")
	}

	zlog.Info().Msgf("tunnel state: %s", tunnelState.String())

	if procs := excludedProcesses.GetProcesses(); procs != nil && len(procs) > 0 {
		for _, proc := range procs {
			zlog.Info().
				Uint32("PID", proc.GetPid()).
				Str("image", proc.GetImage()).
				Bool("inherited", proc.GetInherited()).
				Send()
		}
	} else {
		zlog.Debug().Msg("no excluded processes")
	}

	settings, err := client.mgmtClient.GetSettings(ctx, &emptypb.Empty{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to get settings")
	}

	zlog.Info().Msgf("settings: %s", settings.String())

	lines, err := rl(ctx)
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to initialize readline")
	}

waitLoop:
	for {
		select {
		case sig := <-sigChan:
			zlog.Warn().Msgf("received signal: %s, shutting down mully client", sig.String())
			cancel()
			break waitLoop
		case line := <-lines:
			switch strings.ToLower(line) {
			case "exit":
				cancel()
				break waitLoop
			case "disconnect":
				var pbok = &wrapperspb.BoolValue{}
				if pbok, err = client.mgmtClient.DisconnectTunnel(ctx, &emptypb.Empty{}); err != nil {
					zlog.Error().Err(err).Msg("failed to disconnect tunnel")
				}
				switch {
				case err != nil:
					zlog.Error().Err(err).Msg("failed to disconnect tunnel")
				case pbok.GetValue():
					zlog.Info().Msg("tunnel disconnected")
				case !pbok.GetValue():
					zlog.Warn().Msg("tunnel already disconnected")
				}
			case "connect":
				var pbok = &wrapperspb.BoolValue{}
				if pbok, err = client.mgmtClient.ConnectTunnel(ctx, &emptypb.Empty{}); err != nil {
					zlog.Error().Err(err).Msg("failed to connect tunnel")
				}
				switch {
				case err != nil:
					zlog.Error().Err(err).Msg("failed to connect tunnel")
				case pbok.GetValue():
					zlog.Info().Msg("tunnel connected")
				case !pbok.GetValue():
					zlog.Error().Msg("tunnel already connected")
				}
			}
		}
	}
}
