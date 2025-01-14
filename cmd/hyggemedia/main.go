package main

import (
	"log"
	"os"

	"hyggemedia/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
