package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Azaliya1995/music_library/internal/config"
	"github.com/Azaliya1995/music_library/pkg/log"
	"github.com/Azaliya1995/music_library/version"
)

var serveCmd = &cobra.Command{
	Use:           "serve",
	Short:         "Music Library server",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("Starting server... ")

		cfg, err := config.Init()
		if err != nil {
			log.Error("Init config failed",
				zap.Error(err),
			)
			return err
		}

		logger, err := log.NewLogger(&cfg.LogConfig)
		if err != nil {
			log.Error("Init logger failed",
				zap.Error(err),
			)

			return errors.Wrap(err, "failed to init logger")
		}

		log.SetDefaultLogger(logger)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go waitSignalExit(cancel)

		log.Info("Started server... ")
		<-ctx.Done()

		return nil
	},
}
