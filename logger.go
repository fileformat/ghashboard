package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func initLogger() {

	logLevel := viper.GetString("log-level")
	lvl := slog.LevelInfo
	var err error
	if logLevel != "" {
		ilvl, atoiErr := strconv.Atoi(logLevel)
		if atoiErr == nil {
			lvl = slog.Level(ilvl)
		} else {
			err = lvl.UnmarshalText([]byte(logLevel))
			if err != nil {
				lvl = slog.LevelInfo
			}
		}
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:     lvl,
		AddSource: true,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	if err != nil {
		slog.Error("unable to set log level", "error", err, "level", logLevel)
	} else {
		slog.Debug("log level set", "level", lvl)
	}

}
