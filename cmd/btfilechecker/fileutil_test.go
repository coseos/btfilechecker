// fileutil_test.go
package main

import (
    "os"
    "testing"
)

func TestIsValidFolder(t *testing.T) {
    tmpDir := t.TempDir()
    if !IsValidFolder(tmpDir) {
        t.Errorf("Expected %s to be a valid folder", tmpDir)
    }
    // Test with a file
    tmpFile := tmpDir + "/testfile"
    f, err := os.Create(tmpFile)
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    f.Close()
    if IsValidFolder(tmpFile) {
        t.Errorf("Expected %s to not be a valid folder", tmpFile)
    }
}

