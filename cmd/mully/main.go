package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"git.tcp.direct/kayos/zwrap"
	"github.com/rs/zerolog"
	"golang.org/x/term"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"git.tcp.direct/kayos/mully/pkg/mullvad_mgmt"
	"git.tcp.direct/kayos/mully/pkg/socket"
)

func rl(ctx context.Context) (linrd chan string, promptLock *sync.Mutex, err error) {
	fd := int(os.Stdin.Fd())
	w, h, e := term.GetSize(fd)
	if e != nil {
		return nil, nil, e
	}
	winch := make(chan os.Signal, 1)
	signal.Notify(winch, syscall.SIGWINCH)
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, nil, err
	}
	t := term.NewTerminal(os.Stdin, "mully> ")
	if err = t.SetSize(w, h); err != nil {
		return nil, nil, err
	}

	lines := make(chan string)
	promptLock = &sync.Mutex{}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				promptLock.Lock()
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

	return lines, promptLock, nil
}

var (
	pong = []byte("PONG\n")
	bye  = []byte("BYE\n")
)

func (r *RPCClient) handleSocketCommand(in []byte) (out []byte) {
	switch string(bytes.Fields(in)[0]) {
	case "PING":
		return pong
	case "QUIT":
		go func() {
			time.Sleep(1 * time.Second)
			_ = r.unixS.Close()
		}()
		return bye
	case "RELAY":
		if len(bytes.Fields(in)) < 2 {
			tState, err := r.mgmtClient.GetTunnelState(context.Background(), &emptypb.Empty{})
			if err != nil {
				r.log.Logger.Error().Err(err).Msg("failed to get tunnel state")
				return []byte("ERROR\n")
			}
			connState := tState.GetConnected()
			switch {
			case connState == nil:
				return []byte("DISCONNECTED\n")
			default:
				//
			}
			relayInfo := connState.GetRelayInfo()
			if relayInfo == nil {
				return []byte("DISCONNECTED\n")
			}
			return []byte("RELAY " + relayInfo.TunnelEndpoint.String() + "\n")
		}
		switch string(bytes.Fields(in)[1]) {
		case "RANDOM":
			targetCountry := ""
			if len(bytes.Fields(in)) > 2 {
				targetCountry = string(bytes.Fields(in)[2])
				if len(targetCountry) < 2 {
					return []byte("ERROR\n")
				}

			}
		}
	}
	return nil
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

	lines, plock, err := rl(ctx)
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to initialize readline")
	}

	for {
		select {
		case sig := <-sigChan:
			zlog.Warn().Msgf("received signal: %s, shutting down mully client", sig.String())
			cancel()
			os.Exit(0)
		case line := <-lines:
			switch strings.ToLower(line) {
			case "exit":
				cancel()
				return
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
			case "listen_unix":
				u, e := socket.NewUnixSocket("/tmp/mully.sock", zwrap.Wrap(*zlog))
				if e != nil {
					zlog.Error().Err(e).Msg("failed to create unix socket")
					break
				}
				u.SetCommandHandler(client.handleSocketCommand)
				zlog.Info().Msg("listening on /tmp/mully.sock")
			case "relay":
				var relayInfo *mullvad_mgmt.TunnelEndpoint
				var tState *mullvad_mgmt.TunnelState
				if tState, err = client.mgmtClient.GetTunnelState(ctx, &emptypb.Empty{}); err != nil {
					zlog.Error().Err(err).Msg("failed to get tunnel state")
					break
				}
				relayInfo = tState.GetConnected().GetRelayInfo().GetTunnelEndpoint()
				if relayInfo == nil {
					zlog.Error().Msg("relay info is nil")
					break
				}
				zlog.Info().Msgf("relay: %s", relayInfo.String())

			}
			plock.Unlock()
		}
	}
}
