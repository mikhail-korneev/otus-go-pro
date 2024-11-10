package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TO_PATH = "/tmp/out.txt"
)

var goldenHashes = map[string]string{
	"out_offset0_limit0.txt":       "55aedc99815a6ba613c29881bad22313ad1b067d3b4e24a65a19a2638d874e57",
	"out_offset0_limit10.txt":      "b62c15db701cdb733feb81799bce792338ce56412a146298a23d3b3bf22fcb86",
	"out_offset0_limit1000.txt":    "0a172ebdb46943b4721b7106a55adcaabc40e905602ead85f5075f650394bd5a",
	"out_offset0_limit10000.txt":   "55aedc99815a6ba613c29881bad22313ad1b067d3b4e24a65a19a2638d874e57",
	"out_offset100_limit1000.txt":  "ed523fc33a33551f7578b51df8ad90d85fb82820456330fb7cda166737dc205d",
	"out_offset6000_limit1000.txt": "b21a7efbff10632d261d63c55e475cc55c64d0ef9671e2a046faebd6227e0bbb",
}

func TestCopy(t *testing.T) {

	t.Run("out_offset0_limit0", func(t *testing.T) {

		defer cleanup()

		fromPath := "./testdata/input.txt"
		toPath := TO_PATH
		goldenFile := "out_offset0_limit0.txt"

		offset := int64(0)
		limit := int64(0)

		err := Copy(fromPath, toPath, offset, limit)

		require.NoError(t, err)
		actualHash, err := fileHash(toPath)
		require.NoError(t, err)
		require.Equal(t, goldenHashes[goldenFile], actualHash)
	})

	t.Run("out_offset0_limit10", func(t *testing.T) {

		defer cleanup()

		fromPath := "./testdata/input.txt"
		toPath := TO_PATH
		goldenFile := "out_offset0_limit10.txt"

		offset := int64(0)
		limit := int64(10)

		err := Copy(fromPath, toPath, offset, limit)

		require.NoError(t, err)
		actualHash, err := fileHash(toPath)
		require.NoError(t, err)
		require.Equal(t, goldenHashes[goldenFile], actualHash)
	})

	t.Run("out_offset0_limit1000", func(t *testing.T) {

		defer cleanup()

		fromPath := "./testdata/input.txt"
		toPath := TO_PATH
		goldenFile := "out_offset0_limit1000.txt"

		offset := int64(0)
		limit := int64(1000)

		err := Copy(fromPath, toPath, offset, limit)

		require.NoError(t, err)
		actualHash, err := fileHash(toPath)
		require.NoError(t, err)
		require.Equal(t, goldenHashes[goldenFile], actualHash)
	})

	t.Run("out_offset0_limit10000", func(t *testing.T) {

		defer cleanup()

		fromPath := "./testdata/input.txt"
		toPath := TO_PATH
		goldenFile := "out_offset0_limit10000.txt"

		offset := int64(0)
		limit := int64(10000)

		err := Copy(fromPath, toPath, offset, limit)

		require.NoError(t, err)
		actualHash, err := fileHash(toPath)
		require.NoError(t, err)
		require.Equal(t, goldenHashes[goldenFile], actualHash)
	})

	t.Run("out_offset100_limit1000", func(t *testing.T) {

		defer cleanup()

		fromPath := "./testdata/input.txt"
		toPath := TO_PATH
		goldenFile := "out_offset100_limit1000.txt"

		offset := int64(100)
		limit := int64(1000)

		err := Copy(fromPath, toPath, offset, limit)

		require.NoError(t, err)
		actualHash, err := fileHash(toPath)
		require.NoError(t, err)
		require.Equal(t, goldenHashes[goldenFile], actualHash)
	})

	t.Run("out_offset6000_limit1000", func(t *testing.T) {

		defer cleanup()

		fromPath := "./testdata/input.txt"
		toPath := TO_PATH
		goldenFile := "out_offset6000_limit1000.txt"

		offset := int64(6000)
		limit := int64(1000)

		err := Copy(fromPath, toPath, offset, limit)

		require.NoError(t, err)
		actualHash, err := fileHash(toPath)
		require.NoError(t, err)
		require.Equal(t, goldenHashes[goldenFile], actualHash)
	})
}

func cleanup() {
	if err := removeFile(TO_PATH); err != nil {
		log.Fatalf("failed to remove file with err %s\n", err)
	}
}
