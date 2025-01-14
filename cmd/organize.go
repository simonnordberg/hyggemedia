package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"log/slog"

	"github.com/spf13/cobra"
)

var (
	srcDir  string
	destDir string
)

var validExtensions = []string{".mp4", ".mkv", ".srt"}

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize media files",
	Long:  `Organize media files into a structured format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := organizeFiles(srcDir, destDir, title, dryRun); err != nil {
			slog.Error("Error organizing files", "error", err)
		}
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
	seasonMap, err := scanFiles(srcDir, title)
	if err != nil {
		return err
	}

	for season, files := range seasonMap {
		if err := processSeasonFiles(season, files, destDir, title, dryRun); err != nil {
			return err
		}
	}

	if err := printSummary(seasonMap, title); err != nil {
		return err
	}
	return nil
}

func printSummary(seasonMap map[int][]string, title string) error {
	seasons := make([]int, 0, len(seasonMap))
	for season := range seasonMap {
		seasons = append(seasons, season)
	}
	sort.Ints(seasons)

	fmt.Printf("Organized %d seasons of %s\n", len(seasonMap), title)
	for _, season := range seasons {
		fmt.Printf("Season %d: %d episodes\n", season, len(seasonMap[season]))
	}

	return nil
}

func scanFiles(srcDir, title string) (map[int][]string, error) {
	seasonMap := make(map[int][]string)
	titlePattern := strings.Join(strings.Fields(title), `[\s\.\-_]`)
	re := regexp.MustCompile(`(?i)` + titlePattern + `.*S(\d{1,2})E(\d{1,2})`)

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			slog.Error("Error accessing path", "path", path, "error", err)
			return err
		}
		if !info.IsDir() && isValidExtension(info.Name()) {
			matches := re.FindStringSubmatch(info.Name())
			if len(matches) == 3 {
				season, err := strconv.Atoi(matches[1])
				if err == nil {
					seasonMap[season] = append(seasonMap[season], path)
				} else {
					slog.Warn("Skipping file", "file", path, "reason", err)
				}
			}
		}
		return nil
	})

	return seasonMap, err
}

func processSeasonFiles(season int, files []string, destDir, title string, dryRun bool) error {
	seasonDir := filepath.Join(destDir, fmt.Sprintf("%s/Season %d", title, season))
	if err := createSeasonDir(seasonDir, dryRun); err != nil {
		return err
	}

	for _, file := range files {
		if err := processFile(file, seasonDir, title, dryRun); err != nil {
			slog.Error("Error processing file", "file", file, "reason", err)
			return err
		}
	}

	return nil
}

func createSeasonDir(seasonDir string, dryRun bool) error {
	if dryRun {
		slog.Info("Would create directory", "directory", seasonDir)
		return nil
	}
	if err := os.MkdirAll(seasonDir, os.ModePerm); err != nil {
		return err
	}
	slog.Debug("Created directory", "directory", seasonDir)
	return nil
}

func processFile(file, seasonDir, title string, dryRun bool) error {
	newName, err := generateNewName(filepath.Base(file), title)
	if err != nil {
		slog.Warn("Skipping file", "file", file, "reason", err)
		return err
	}
	destPath := filepath.Join(seasonDir, newName)
	if dryRun {
		slog.Info("Would copy file", "source", file, "destination", destPath)
		return nil
	}
	if err := copyFile(file, destPath); err != nil {
		slog.Error("Error copying file", "source", file, "destination", destPath, "reason", err)
		return err
	}
	slog.Debug("Copied file", "source", file, "destination", destPath)
	return nil
}

func generateNewName(oldName, showName string) (string, error) {
	re := regexp.MustCompile(`(?i)S(\d{1,2})E(\d{1,2})`)
	matches := re.FindStringSubmatch(oldName)
	if len(matches) != 3 {
		return "", fmt.Errorf("could not extract season and episode information from %s", oldName)
	}

	season := matches[1]
	episode := matches[2]

	newName := fmt.Sprintf("%s S%sE%s%s", showName, season, episode, filepath.Ext(oldName))
	return newName, nil
}

func isValidExtension(fileName string) bool {
	for _, ext := range validExtensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}
	return false
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

	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return destFile.Sync()
}
