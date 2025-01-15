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

type MovieOrganizer struct{}

func (o MovieOrganizer) Organize(title, srcDir, destDir string, dryRun, move bool) error {
	err := filepath.Walk(srcDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !utils.IsMediaFile(path) {
			return nil
		}

		info, err := parseMovieInfo(title, path)
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

func parseMovieInfo(title, file string) (MovieInfo, error) {
	titlePattern := strings.Join(strings.Fields(title), `[\s\.\-_]`)
	re := regexp.MustCompile(`(?i)` + titlePattern + `.*[\s\.\-_]+(\d{4})[\s\.\-_]+`)
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

func (e MovieInfo) DestinationFilename() string {
	return fmt.Sprintf("%s (%04d)%s", e.Title, e.Year, filepath.Ext(e.Filename))
}

func (e MovieInfo) DestinationDirname() string {
	return fmt.Sprintf("%s (%04d)", e.Title, e.Year)
}
