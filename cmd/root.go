package main

import (
	"fmt"

	"github.com/Azaliya1995/music_library/version"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "",
		Long:          "Music Library: " + version.Version,
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       fmt.Sprintf("%s %s", version.Version, version.CommitHash),
	}

	cmd.AddCommand(serveCmd)
	cmd.AddCommand(migrateCMD)

	return cmd
}

func Execute() (err error) {
	return rootCmd().Execute()
}
