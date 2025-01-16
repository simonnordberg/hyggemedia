package cmd

import (
	"hyggemedia/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var cfg = &config.Config{}

var rootCmd = &cobra.Command{
	Use:   "hyggemedia",
	Short: "organize media files",
	Long:  `Hygge Media is a command line application designed to rename and organize media files in a directory to match the format prescribed by Emby. This tool simplifies the organization of your media library, ensuring that your files are named consistently and correctly.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Title, "title", "", "Title of the TV show or movie")
	rootCmd.MarkPersistentFlagRequired("title")

	rootCmd.PersistentFlags().StringVar(&cfg.SourceDir, "source-dir", "", "Source directory to scan for files")
	rootCmd.MarkPersistentFlagRequired("source-dir")
	rootCmd.MarkPersistentFlagDirname("source-dir")

	rootCmd.PersistentFlags().StringVar(&cfg.TargetDir, "target-dir", "", "Target directory to organize files into")
	rootCmd.MarkPersistentFlagRequired("target-dir")
	rootCmd.MarkPersistentFlagDirname("target-dir")

	rootCmd.PersistentFlags().BoolVar(&cfg.Exec, "exec", false, "Execute the changes")
	rootCmd.PersistentFlags().BoolVar(&cfg.Move, "move", false, "Move files instead of copying")
}
