package main

import (
	"hyggemedia/media"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var title, srcDir, destDir string
var dryRun, move bool

var tvCmd = &cobra.Command{
	Use:   "tv",
	Short: "Organize TV shows",
	Run: func(cmd *cobra.Command, args []string) {
		media.TVOrganizer{}.Organize(title, srcDir, destDir, dryRun, move)
	},
}

var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "Organize movies",
	Run: func(cmd *cobra.Command, args []string) {
		media.MovieOrganizer{}.Organize(title, srcDir, destDir, dryRun, move)
	},
}

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.PersistentFlags().StringVar(&title, "title", "", "Title of the TV show or movie")
	rootCmd.PersistentFlags().StringVar(&srcDir, "src-dir", "", "Source directory to scan for files")
	rootCmd.PersistentFlags().StringVar(&destDir, "dest-dir", "", "Destination directory to organize files into")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", true, "Perform a dry run without making changes")
	rootCmd.PersistentFlags().BoolVar(&move, "move", false, "Move files instead of copying")
	rootCmd.MarkPersistentFlagRequired("dest-dir")
	rootCmd.MarkPersistentFlagRequired("src-dir")
	rootCmd.MarkPersistentFlagRequired("title")
	rootCmd.MarkPersistentFlagDirname("dest-dir")
	rootCmd.MarkPersistentFlagDirname("src-dir")

	rootCmd.AddCommand(tvCmd)
	rootCmd.AddCommand(movieCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
