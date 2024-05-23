package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Account  string
	logLevel zerolog.Level
	timeout  time.Duration
}

func flagConfigVals(conf *Config) error {
	badArg := func(arg string) error {
		return fmt.Errorf("unknown argument: %s", arg)
	}

	verbosity := func(arg string) (error, bool) {
		if !strings.HasPrefix(arg, "-v") {
			return nil, true
		}
		if strings.Count(arg, "v") == 1 {
			conf.logLevel = zerolog.DebugLevel
			return nil, false
		}
		for _, c := range arg[1:] {
			if c != 'v' {
				return badArg(arg), false
			}
		}
		conf.logLevel = zerolog.TraceLevel
		return nil, false
	}

	for _, arg := range os.Args[1:] {
		err, keepGoing := verbosity(arg)
		if err != nil {
			return err
		}
		if !keepGoing {
			continue
		}

		switch arg {
		case "--debug":
			conf.logLevel = zerolog.DebugLevel
		case "--trace":
			conf.logLevel = zerolog.TraceLevel
		default:
		}

		switch {
		case conf.Account != "":
			return badArg(arg)
		case conf.Account == "" && !strings.HasPrefix(arg, "-"):
			conf.Account = arg
		case conf.Account == "" && strings.HasPrefix(arg, "-"):
			return badArg(arg)
		default:
		}
	}

	return nil
}

func getConfig() (*Config, error) {
	if len(os.Args) < 2 && os.Getenv("MULLVAD_ACCOUNT") == "" {
		return nil, fmt.Errorf("missing mullvad account key")
	}

	conf := &Config{
		logLevel: zerolog.InfoLevel,
		timeout:  5 * time.Second,
	}
	errs := make([]error, 0, 2)
	errs = append(errs, flagConfigVals(conf))

	if os.Getenv("MULLVAD_ACCOUNT") != "" && conf.Account == "" {
		conf.Account = os.Getenv("MULLVAD_ACCOUNT")
	}

	if conf.Account == "" {
		errs = append(errs, fmt.Errorf("missing mullvad account key"))
	}

	return conf, errors.Join(errs...)
}
