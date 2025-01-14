package cmd

import (
	"github.com/spf13/cobra"
)

var (
	title  string
	dryRun bool
)

var rootCmd = &cobra.Command{
	Use:   "hyggemedia",
	Short: "Hygge Media CLI",
	Long:  `Hygge Media is a CLI application for managing media files.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "Perform a dry run without making changes")
	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Title of the show (mandatory)")
	rootCmd.MarkPersistentFlagRequired("title")
}
