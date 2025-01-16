package report

import (
	"fmt"
	"hyggemedia/internal/config"
	"hyggemedia/internal/file"
)

func Report(config *config.Config, changes file.Changes) {
	fmt.Println("Changes:")
	for _, change := range changes {
		fmt.Printf("Source: %s\n", change.Source)
		fmt.Printf("Target: %s\n", change.Target)
		fmt.Println()
	}
}
