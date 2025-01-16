package rename

import (
	"fmt"
	"hyggemedia/internal/config"
	"hyggemedia/internal/file"
	"hyggemedia/internal/utils"
	"path/filepath"
)

func Rename(config *config.Config, changes file.Changes) error {
	for _, change := range changes {
		destDir := filepath.Dir(change.Target)
		if err := utils.CreateDir(destDir); err != nil {
			fmt.Println("Error creating directory", destDir)
			continue
		}

		if err := utils.MoveOrCopyFile(change.Source, change.Target, config.Move); err != nil {
			fmt.Println("Error moving", change.Source, "to", change.Target)
			continue
		}
	}
	return nil
}
