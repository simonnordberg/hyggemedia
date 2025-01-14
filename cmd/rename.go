package cmd

import (
	"fmt"
	"log"

	"hyggemedia/internal/rename"

	"github.com/spf13/cobra"
)

var (
	dir string
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename media files",
	Long:  `Rename media files to match the Emby format.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		if err := rename.RenameFiles(dir, dryRun, title); err != nil {
			log.Fatalf("Error renaming files: %v", err)
		}
		fmt.Println("Files renamed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringVarP(&dir, "dir", "d", "", "Directory to rename files (mandatory)")
	renameCmd.MarkFlagRequired("show-name")
	renameCmd.MarkFlagRequired("dir")
}
