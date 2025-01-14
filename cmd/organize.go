package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize media files",
	Long:  `Organize media files into a structured format.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		// Implement the organize functionality here
		fmt.Printf("Organizing files in directory: %s\n", dir)
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
}
