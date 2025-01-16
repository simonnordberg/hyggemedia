package find

import (
	"hyggemedia/internal/config"
	"hyggemedia/internal/file"
	"os"
	"path/filepath"
	"strings"
)

type MediaFinder interface {
	Find(conf *config.Config) (file.Changes, error)
	ParseMediaInfo(title, file string) (MediaInfo, error)
}

type MediaInfo interface {
	DestFilename() string
	DestDirname() string
}

func isMediaFile(file string) bool {
	mediaExtensions := []string{".mp4", ".mkv", ".avi", ".srt"}
	for _, ext := range mediaExtensions {
		if strings.HasSuffix(strings.ToLower(file), ext) {
			return true
		}
	}
	return false
}

func Find(finder MediaFinder, conf *config.Config) (file.Changes, error) {
	var changes file.Changes
	err := filepath.Walk(conf.SourceDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !isMediaFile(path) {
			return nil
		}

		info, err := finder.ParseMediaInfo(conf.Title, path)
		if err != nil {
			return nil
		}

		destDir := filepath.Join(conf.TargetDir, info.DestDirname())
		destFile := filepath.Join(destDir, info.DestFilename())

		changes = append(changes, &file.Change{
			Source: path,
			Target: destFile,
		})
		return nil
	})
	return changes, err
}
