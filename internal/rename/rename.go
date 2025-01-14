package rename

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// RenameFiles renames files in the specified directory to match the Emby format.
func RenameFiles(dir string, dryRun bool, showName string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if ext != ".mp4" && ext != ".mkv" && ext != ".srt" {
				return nil
			}

			newName, err := GenerateNewName(info.Name(), showName)
			if err != nil {
				log.Printf("Skipping file: %s, reason: %v\n", path, err)
				return nil
			}
			newPath := filepath.Join(filepath.Dir(path), newName)
			if dryRun {
				log.Printf("Would rename: %s -> %s\n", path, newPath)
			} else {
				if err := os.Rename(path, newPath); err != nil {
					log.Printf("Error renaming file: %s -> %s, reason: %v\n", path, newPath, err)
					return err
				}
				log.Printf("Renamed: %s -> %s\n", path, newName)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking the path %q: %v\n", dir, err)
	}
	return err
}

// generateNewName generates the new name according to Emby format.
func GenerateNewName(oldName, showName string) (string, error) {
	// Example regex to extract season and episode information
	re := regexp.MustCompile(`(?i)S(\d{1,2})E(\d{1,2})`)
	matches := re.FindStringSubmatch(oldName)
	if len(matches) != 3 {
		return "", logErrorf("could not extract season and episode information from %s", oldName)
	}

	season := matches[1]
	episode := matches[2]

	// Format the new name
	newName := fmt.Sprintf("%s S%sE%s%s", showName, season, episode, filepath.Ext(oldName))
	return newName, nil
}

// logErrorf logs an error message and returns an error.
func logErrorf(format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	log.Println(err)
	return err
}
