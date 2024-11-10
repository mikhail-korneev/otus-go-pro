package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func fileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf("%x", hash.Sum(nil))
	return res, nil
}

func removeFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		fmt.Printf("failed to remove file with err %s\n", err)
		return err
	}
	return nil
}
