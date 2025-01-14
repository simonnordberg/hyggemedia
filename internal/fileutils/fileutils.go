package fileutils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func CreateDir(dir string, dryRun bool) error {
	if dryRun {
		logrus.Info("Would create directory", "directory", dir)
		return nil
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	logrus.Debug("Created directory", "directory", dir)
	return nil
}

func MoveFile(src, dst string, dryRun bool) error {
	if dryRun {
		logrus.Info("Would move file", "source", src, "destination", dst)
		return nil
	}
	return os.Rename(src, dst)
}

func CopyFile(src, dst string, dryRun bool) error {
	if dryRun {
		logrus.Info("Would copy file", "source", src, "destination", dst)
		return nil
	}
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return destFile.Sync()
}
