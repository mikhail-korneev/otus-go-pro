package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		fmt.Printf("failed with err %s\n", err)
		return err
	}
	defer src.Close()

	fromInfo, err := src.Stat()
	if err != nil {
		fmt.Printf("failed with err %s\n", err)
		return err
	}

	if err := checkPreconditions(fromInfo); err != nil {
		fmt.Println("check preconditions failed")
		return err
	}

	dst, err := os.Create(toPath)
	if err != nil {
		fmt.Printf("failed with err %s\n", err)
		return err
	}
	defer dst.Close()

	if limit == 0 {
		limit = fromInfo.Size()
	}
	bar := pb.Full.Start64(limit)
	sectionReader := io.NewSectionReader(src, offset, limit)
	barReader := bar.NewProxyReader(sectionReader)
	io.Copy(dst, barReader)
	bar.Finish()

	return nil
}

func checkPreconditions(info os.FileInfo) error {
	if info.IsDir() {
		fmt.Printf("check failed: %s is directory\n", info.Name())
		return ErrUnsupportedFile
	}
	if info.Mode()&os.ModeDevice != 0 {
		fmt.Printf("check failed: %s is pseudo-file\n", info.Name())
		return ErrUnsupportedFile
	}
	if offset > info.Size() {
		fmt.Printf("check failed: offset %d > file size %d\n", offset, info.Size())
		return ErrOffsetExceedsFileSize
	}
	return nil
}
