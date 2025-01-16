package utils

import (
	"fmt"
	"io"
	"os"
)

func MoveOrCopyFile(src, dst string, move bool) error {
	if move {
		return os.Rename(src, dst)
	}
	return CopyFile(src, dst)
}

func CopyFile(src, dst string) error {
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

func CreateDir(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	fmt.Println("Created directory", dir)
	return nil
}
