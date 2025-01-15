package media

import (
	"hyggemedia/utils"
	"os"
	"path/filepath"
	"strings"
)

type MediaOrganizer interface {
	Organize(title, srcDir, destDir string, dryRun, move bool) error
	ParseMediaInfo(title, file string) (MediaInfo, error)
}

type MediaInfo interface {
	DestFilename() string
	DestDirname() string
}

func Organize(o MediaOrganizer, title, srcDir, destDir string, dryRun, move bool) error {
	err := filepath.Walk(srcDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !isMediaFile(path) {
			return nil
		}

		info, err := o.ParseMediaInfo(title, path)
		if err != nil {
			return nil
		}

		destDir := filepath.Join(destDir, info.DestDirname())
		if err := utils.CreateDir(destDir, dryRun); err != nil {
			return err
		}

		destFile := filepath.Join(destDir, info.DestFilename())
		if err := utils.MoveOrCopyFile(path, destFile, move, dryRun); err != nil {
			return err
		}
		return nil
	})
	return err
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
