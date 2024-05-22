package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/reflect/protoreflect"

	"git.tcp.direct/kayos/mully/pkg/mullvad_mgmt"
)

var mProto = mullvad_mgmt.File_management_interface_proto

func checkProto(debug bool) protoreflect.ServiceDescriptor {
	if mProto.Services().Len() == 0 {
		panic("mullvad management interface proto has no services, did we forget to generate from pb?")
	}
	if mProto.Services().Len() > 1 {
		println(strings.Repeat("=", 80))
		for n := 0; n < mProto.Services().Len(); n++ {
			spew.Dump(mProto.Services().Get(n))
		}
		print(strings.Repeat("=", 80))
		panic("mullvad management interface proto has more than one service...!?")
	}
	if mProto.Services().Get(0).Methods().Len() == 0 {
		panic("mullvad management interface proto has no methods, did we forget to generate from pb?")
	}
	if debug {
		log := initConsoleLogger().WithForceLevel(zerolog.DebugLevel)
		slog := log.Logger.With().Str("caller", string(mProto.Services().Get(0).Name())).Logger()
		for n := 0; n < mProto.Services().Get(0).Methods().Len(); n++ {
			slog.Debug().Msgf("method %d: %s", n, mProto.Services().Get(0).Methods().Get(n).Name())
		}
	}
	return mProto.Services().Get(0)
}
