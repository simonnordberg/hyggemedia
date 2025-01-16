package cmd

import (
	"hyggemedia/internal/find"
	"hyggemedia/internal/handlers"

	"github.com/spf13/cobra"
)

var tvCmd = &cobra.Command{
	Use:   "tv",
	Short: "Organize TV shows",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Organize(cfg, find.TVMediaFinder{})
	},
}

func init() {
	rootCmd.AddCommand(tvCmd)
}