package fileutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	testDir := filepath.Join(dir, "newdir")

	// Test dry run
	err = CreateDir(testDir, true)
	assert.NoError(t, err)
	_, err = os.Stat(testDir)
	assert.True(t, os.IsNotExist(err))

	// Test actual creation
	err = CreateDir(testDir, false)
	assert.NoError(t, err)
	_, err = os.Stat(testDir)
	assert.False(t, os.IsNotExist(err))
}

func TestMoveFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	srcFile := filepath.Join(dir, "src.txt")
	dstFile := filepath.Join(dir, "dst.txt")

	err = os.WriteFile(srcFile, []byte("test content"), 0644)
	assert.NoError(t, err)

	// Test dry run
	err = MoveFile(srcFile, dstFile, true)
	assert.NoError(t, err)
	_, err = os.Stat(srcFile)
	assert.False(t, os.IsNotExist(err))
	_, err = os.Stat(dstFile)
	assert.True(t, os.IsNotExist(err))

	// Test actual move
	err = MoveFile(srcFile, dstFile, false)
	assert.NoError(t, err)
	_, err = os.Stat(srcFile)
	assert.True(t, os.IsNotExist(err))
	_, err = os.Stat(dstFile)
	assert.False(t, os.IsNotExist(err))
}

func TestCopyFile(t *testing.T) {
	dir, err := os.MkdirTemp("", "testdir")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	srcFile := filepath.Join(dir, "src.txt")
	dstFile := filepath.Join(dir, "dst.txt")

	err = os.WriteFile(srcFile, []byte("test content"), 0644)
	assert.NoError(t, err)

	// Test dry run
	err = CopyFile(srcFile, dstFile, true)
	assert.NoError(t, err)
	_, err = os.Stat(dstFile)
	assert.True(t, os.IsNotExist(err))

	// Test actual copy
	err = CopyFile(srcFile, dstFile, false)
	assert.NoError(t, err)
	_, err = os.Stat(dstFile)
	assert.False(t, os.IsNotExist(err))

	content, err := os.ReadFile(dstFile)
	assert.NoError(t, err)
	assert.Equal(t, "test content", string(content))
}
