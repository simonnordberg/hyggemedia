package cmd

import (
	"github.com/spf13/cobra"
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
	rootCmd.PersistentFlags().StringP("dir", "d", "", "Directory to operate on (mandatory)")
	rootCmd.MarkPersistentFlagRequired("dir")
}
