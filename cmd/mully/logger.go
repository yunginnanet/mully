package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"git.tcp.direct/kayos/zwrap"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/connectivity"
)

var stateColors = map[string]int{
	"CONNECTING":        33,
	"UNKNOWN":           90,
	"DEBUG":             90,
	"READY":             32,
	"IDLE":              37,
	"TRANSIENT_FAILURE": 31,
	"SHUTDOWN":          31,
	"STATE":             36,
}

var replacer = strings.NewReplacer(
	"]", "",
	"[", "",
	"#", "",
)

func debugFormatPrep(fields map[string]interface{}) error {
	fields["state"] = "DEBUG"
	if _, ok := fields["caller"]; !ok {
		fields["caller"] = "gRPC_internal"
		fields["message"] = strings.Replace(fields["message"].(string), "[core]", "", 1)
	}
	fields["message"] = strings.ReplaceAll(fields["message"].(string), "]", "] ")
	msgString := fields["message"].(string)
	msgFields := strings.Fields(msgString)

	remove := func(i int) {
		fields["message"] = strings.TrimSpace(
			strings.ReplaceAll(
				fields["message"].(string),
				strings.Join(msgFields[i:i+2], " "),
				"",
			),
		)
	}

	for i := range msgFields {
		tmp := replacer.Replace(msgFields[i])
		_, hasChannel := fields["channel"]
		_, hasSubChannel := fields["sub_channel"]
		if !hasChannel && strings.EqualFold(tmp, "channel") && len(msgFields) > i+1 {
			fields["channel"] = replacer.Replace(msgFields[i+1])
			remove(i)
			continue
		}
		if !hasSubChannel && strings.EqualFold(tmp, "subchannel") && len(msgFields) > i+1 {
			fields["sub_channel"] = replacer.Replace(msgFields[i+1])
			remove(i)
			continue
		}
	}

	return nil
}

func checkTransitionState(fields map[string]interface{}) {
	if _, ok := fields["message"]; !ok {
		return
	}
	if fields["message"].(string) != "state changed" {
		return
	}
	if st, ok := fields["state"]; !ok || st.(string) == "" {
		return
	}
	fields["message"] = fields["state"].(string)
	fields["state"] = "STATE"
}

func formatPrep(fields map[string]interface{}) error {
	checkTransitionState(fields)
	if _, ok := fields["state"]; !ok {
		if err := debugFormatPrep(fields); err != nil {
			panic(err)
		}
	}

	colInt, ok := stateColors[fields["state"].(string)]
	if !ok {
		colInt = 90
	}

	stateStrTrunc := fields["state"].(string)
	if stateStrTrunc == "CONNECTING" {
		stateStrTrunc = "START"
	}
	if len(stateStrTrunc) > 5 {
		stateStrTrunc = stateStrTrunc[:5]
	}

	fields["state"] = fmt.Sprintf("[%s]",
		zwrap.Colorize(stateStrTrunc, colInt, runtime.GOOS == "windows"),
	)

	return nil
}

func initConsoleLogger(hook *logHooker, cs ...*Config) *zwrap.Logger {
	if len(cs) > 0 {
		zerolog.SetGlobalLevel(cs[0].logLevel)
	}

	// []string{"time", "level", "caller", "message"},

	zlc := zerolog.ConsoleWriter{
		Out:           os.Stderr,
		FormatLevel:   zwrap.LogLevelFmt(runtime.GOOS == "windows"),
		PartsOrder:    []string{"time", "level", "state", "caller", "message"},
		FieldsExclude: []string{"state"},
		FormatPrepare: formatPrep,
	}

	zl := zerolog.New(zlc).With().Timestamp().Logger()

	if hook != nil {
		zl = zerolog.New(zlc).With().Timestamp().Logger()
		zl = zl.Hook(hook)
	}

	return zwrap.Wrap(zl)
}

// because there's no need to export `Run`, this is a hack
type logHooker struct {
	rpc *RPCClient
}

func (h *logHooker) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	h.rpc.hookRun(e, level, msg)
}

func (r *RPCClient) hookRun(e *zerolog.Event, level zerolog.Level, msg string) {
	state := r.GetState()
	e.Str("state", state.String())
	if state != connectivity.Ready || msg == "state changed" {
		e.Str("caller", r.CanonicalTarget())
		return
	}
	e.Str("caller", "mullvad-daemon")
}
