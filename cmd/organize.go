package cmd

import (
	"fmt"
	"hyggemedia/internal/rename"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	srcDir  string
	destDir string
)

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize media files",
	Long:  `Organize media files into a structured format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := organizeFiles(srcDir, destDir, title, dryRun); err != nil {
			log.Fatalf("Error organizing files: %v", err)
		}
		fmt.Println("Files organized successfully.")
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
	organizeCmd.Flags().StringVarP(&srcDir, "src-dir", "s", "", "Source directory to scan for files (mandatory)")
	organizeCmd.Flags().StringVarP(&destDir, "dest-dir", "d", "", "Destination directory to organize files into (mandatory)")
	organizeCmd.MarkFlagRequired("src-dir")
	organizeCmd.MarkFlagRequired("dest-dir")
}

func organizeFiles(srcDir, destDir, title string, dryRun bool) error {
	seasonMap := make(map[int][]string)

	titlePattern := strings.Join(strings.Fields(title), `[\s\.\-_]`)
	re := regexp.MustCompile(`(?i)` + titlePattern + `.*S(\d{1,2})E(\d{1,2})`)

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if ext != ".mp4" && ext != ".mkv" && ext != ".srt" {
				return nil
			}

			matches := re.FindStringSubmatch(info.Name())
			if len(matches) != 3 {
				return nil
			}

			season, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Printf("Skipping file: %s, reason: %v\n", path, err)
				return nil
			}

			seasonMap[season] = append(seasonMap[season], path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for season, files := range seasonMap {
		seasonDir := filepath.Join(destDir, fmt.Sprintf("%s/Season %d", title, season))
		if dryRun {
			log.Printf("Would create directory: %s\n", seasonDir)
		} else {
			if err := os.MkdirAll(seasonDir, os.ModePerm); err != nil {
				return err
			}
		}

		for _, file := range files {
			newName, err := rename.GenerateNewName(filepath.Base(file), title)
			if err != nil {
				log.Printf("Skipping file: %s, reason: %v\n", file, err)
				continue
			}

			destPath := filepath.Join(seasonDir, newName)
			if dryRun {
				log.Printf("Would copy: %s -> %s\n", file, destPath)
			} else {
				if err := copyFile(file, destPath); err != nil {
					log.Printf("Error copying file: %s -> %s, reason: %v\n", file, destPath, err)
					return err
				}
				log.Printf("Copied: %s -> %s\n", file, destPath)
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
