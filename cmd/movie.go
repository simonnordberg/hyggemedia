package cmd

import (
	"hyggemedia/internal/find"
	"hyggemedia/internal/handlers"

	"github.com/spf13/cobra"
)

var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "Organize movie files",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.OrganizeMediaFiles(cfg, find.MovieMediaParser{})
	},
}

func init() {
	rootCmd.AddCommand(movieCmd)
}
