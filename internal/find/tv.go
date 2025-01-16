package find

import (
	"fmt"
	"hyggemedia/internal/config"
	"hyggemedia/internal/file"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type TVMediaFinder struct{}

func (f TVMediaFinder) Find(conf *config.Config) (file.Changes, error) {
	return Find(f, conf)
}

func (o TVMediaFinder) ParseMediaInfo(title, file string) (MediaInfo, error) {
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

func (e EpisodeInfo) DestFilename() string {
	return fmt.Sprintf("%s S%02dE%02d%s", e.Title, e.Season, e.Episode, filepath.Ext(e.Filename))
}

func (e EpisodeInfo) DestDirname() string {
	return fmt.Sprintf("%s/Season %d", e.Title, e.Season)
}