package cmd

import (
	"fmt"
	"log"

	"hyggemedia/internal/rename"

	"github.com/spf13/cobra"
)

var (
	showName string
	dryRun   bool
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename media files",
	Long:  `Rename media files to match the Emby format.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		if err := rename.RenameFiles(dir, dryRun, showName); err != nil {
			log.Fatalf("Error renaming files: %v", err)
		}
		fmt.Println("Files renamed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringVarP(&showName, "show-name", "s", "", "Name of the show (mandatory)")
	renameCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false, "Perform a dry run without renaming files")
	renameCmd.MarkFlagRequired("show-name")
}
