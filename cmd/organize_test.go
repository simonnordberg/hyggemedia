package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateNewName(t *testing.T) {
	tests := []struct {
		oldName  string
		showName string
		expected string
	}{
		{"show.S01E01.mkv", "Show", "Show S01E01.mkv"},
		{"show.s02e10.mp4", "Show", "Show S02E10.mp4"},
		{"show.s3e5.srt", "Show", "Show S03E05.srt"},
	}

	for _, test := range tests {
		newName, err := generateFilename(test.oldName, test.showName)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if newName != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, newName)
		}
	}
}

func TestIsValidExtension(t *testing.T) {
	tests := map[string]bool{
		"video.mp4":    true,
		"video.mkv":    true,
		"subtitle.srt": true,
		"document.txt": false,
	}

	for fileName, expected := range tests {
		if isValidExtension(fileName) != expected {
			t.Errorf("Expected %v for %s, got %v", expected, fileName, !expected)
		}
	}
}

func TestScanFiles(t *testing.T) {
	srcDir := t.TempDir()
	title := "Show"
	files := []string{
		"Show.S01E01.mkv",
		"Show.S01E02.mp4",
		"Show.S02E01.srt",
		"Show.S02E02.mkv",
	}

	for _, file := range files {
		if _, err := os.Create(filepath.Join(srcDir, file)); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	seasonMap, err := scanFiles(srcDir, title)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedSeasons := 2
	expectedEpisodes := 2
	if len(seasonMap) != expectedSeasons {
		t.Errorf("Expected %d seasons, got %d", expectedSeasons, len(seasonMap))
	}

	for season, episodes := range seasonMap {
		if len(episodes) != expectedEpisodes {
			t.Errorf("Expected %d episodes in season %d, got %d", expectedEpisodes, season, len(episodes))
		}
	}
}

func TestProcessSeasonFiles(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()
	title := "Show"
	files := []string{
		filepath.Join(srcDir, "Show.S01E01.mkv"),
		filepath.Join(srcDir, "Show.S01E02.mp4"),
	}

	for _, file := range files {
		if _, err := os.Create(file); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	err := processSeasonFiles(1, files, destDir, title, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	seasonDir := filepath.Join(destDir, "Show/Season 1")
	for _, file := range files {
		newName, _ := generateFilename(file, title)
		if _, err := os.Stat(filepath.Join(seasonDir, newName)); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist", filepath.Join(seasonDir, newName))
		}
	}
}
