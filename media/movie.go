package media

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type MovieOrganizer struct{}

func (o MovieOrganizer) Organize(title, srcDir, destDir string, dryRun, move bool) error {
	return Organize(o, title, srcDir, destDir, dryRun, move)
}

func (o MovieOrganizer) ParseMediaInfo(title, file string) (MediaInfo, error) {
	titlePattern := strings.Join(strings.Fields(title), `[\s\.\-_]`)
	re := regexp.MustCompile(`(?i)` + titlePattern + `.*[\s\.\-_\(\)]+(\d{4})[\s\.\-_\(\)]+`)
	if matches := re.FindStringSubmatch(file); len(matches) == 2 {
		year, _ := strconv.Atoi(matches[1])
		return MovieInfo{
			Filename: file,
			Title:    title,
			Year:     year,
		}, nil
	}
	return MovieInfo{}, fmt.Errorf("failed to parse movie info: %s", file)
}

type MovieInfo struct {
	Filename string
	Title    string
	Year     int
}

func (e MovieInfo) DestFilename() string {
	return fmt.Sprintf("%s (%04d)%s", e.Title, e.Year, filepath.Ext(e.Filename))
}

func (e MovieInfo) DestDirname() string {
	return fmt.Sprintf("%s (%04d)", e.Title, e.Year)
}
