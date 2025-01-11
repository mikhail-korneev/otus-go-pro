package main

import (
	"os"
	"testing"
)

func setup(t *testing.T, pattern string) (string, func()) {
	t.Helper()
	t.Log("run before test")

	tmpDir, err := os.MkdirTemp("", pattern)
	if err != nil {
		t.Fatalf("failed to create temp dir with err %s\n", err)
	}
	t.Logf("temp dir created: %s\n", tmpDir)

	return tmpDir, func() {
		t.Log("run after test")
		os.RemoveAll(tmpDir)
		t.Logf("temp dir cleaned up: %s\n", tmpDir)
	}
}
