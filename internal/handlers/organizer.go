package handlers

import (
	"hyggemedia/internal/config"
	"hyggemedia/internal/file"
	"hyggemedia/internal/find"
	"hyggemedia/internal/rename"
	"hyggemedia/internal/report"
)

func applyChanges(config *config.Config, changes file.Changes) error {
	if len(changes) == 0 {
		return nil
	}

	if !config.Exec {
		report.Report(config, changes)
		return nil
	}

	err := rename.Rename(config, changes)
	return err
}

func Organize(config *config.Config, finder find.MediaFinder) error {
	changes, err := find.Find(finder, config)
	if err != nil {
	}

	err = applyChanges(config, changes)
	return err
}
