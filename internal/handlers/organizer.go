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

func OrganizeMediaFiles(config *config.Config, parser find.MediaParser) error {
	changes, err := find.Find(parser, config)
	if err != nil {
		return err
	}

	err = applyChanges(config, changes)
	return err
}
