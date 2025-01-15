package media

import (
	"fmt"
	"hyggemedia/utils"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type TVOrganizer struct{}

func (o TVOrganizer) Organize(title, srcDir, destDir string, dryRun, move bool) error {
	err := filepath.Walk(srcDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !utils.IsMediaFile(path) {
			return nil
		}

		info, err := parseEpisodeInfo(title, path)
		if err != nil {
			return nil
		}

		if err := utils.CreateDir(info.DestinationDirname(), dryRun); err != nil {
			return err
		}

		destFile := filepath.Join(destDir, info.DestinationDirname(), info.DestinationFilename())
		if err := utils.MoveOrCopyFile(path, destFile, move, dryRun); err != nil {
			return err
		}
		return nil
	})
	return err
}

func parseEpisodeInfo(title, file string) (EpisodeInfo, error) {
	titlePattern := strings.Join(strings.Fields(title), `[\s\.\-_]`)
	re := regexp.MustCompile(`(?i)` + titlePattern + `.*S(\d{1,2})E(\d{1,2})`)
	if matches := re.FindStringSubmatch(file); len(matches) == 3 {
		season, _ := strconv.Atoi(matches[1])
		episode, _ := strconv.Atoi(matches[2])
		return EpisodeInfo{
			Filename: file,
			Title:    title,
			Season:   season,
			Episode:  episode,
		}, nil
	}
	return EpisodeInfo{}, fmt.Errorf("failed to parse episode info: %s", file)
}

type EpisodeInfo struct {
	Filename string
	Title    string
	Season   int
	Episode  int
}

func (e EpisodeInfo) DestinationFilename() string {
	return fmt.Sprintf("%s S%02dE%02d%s", e.Title, e.Season, e.Episode, filepath.Ext(e.Filename))
}

func (e EpisodeInfo) DestinationDirname() string {
	return fmt.Sprintf("%s/Season %d", e.Title, e.Season)
}
