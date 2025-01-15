package utils

import (
	"fmt"
	"io"
	"os"
)

func MoveOrCopyFile(src, dst string, move, dryRun bool) error {
	if move {
		return MoveFile(src, dst, dryRun)
	}
	return CopyFile(src, dst, dryRun)
}

func CopyFile(src, dst string, dryRun bool) error {
	if dryRun {
		fmt.Println("Would copy", src, "to", dst)
		return nil
	}
	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err == nil {
		fmt.Println("Copied", src, "to", dst)
	}
	return err
}

func MoveFile(src, dst string, dryRun bool) error {
	if dryRun {
		fmt.Println("Would move", src, "to", dst)
		return nil
	}
	return os.Rename(src, dst)
}

func CreateDir(dir string, dryRun bool) error {
	if dryRun {
		fmt.Println("Would create directory", dir)
		return nil
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	fmt.Println("Created directory", dir)
	return nil
}
