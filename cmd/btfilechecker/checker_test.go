package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBTFileChecker_Check(t *testing.T) {
	// Set up a temporary directory to act as the base path
	tmpDir := t.TempDir()

	// Create a file that should exist
	existingFile := "exists.txt"
	existingFilePath := filepath.Join(tmpDir, existingFile)
	if err := os.WriteFile(existingFilePath, []byte("data"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Entry for the existing file
	entryExists := BTFileEntry{
		Label:    LabelBT,
		Name:     "Test Exists",
		Filename: existingFile,
	}

	// Entry for a missing file
	entryMissing := BTFileEntry{
		Label:    LabelBT,
		Name:     "Test Missing",
		Filename: "missing.txt",
	}

	checker := NewBTFileChecker(tmpDir + "/") // Ensure trailing slash

	// Test existing file
	if err := checker.Check(entryExists); err != nil {
		t.Errorf("Checker reported missing file that exists: %v", err)
	}

	// Test missing file
	if err := checker.Check(entryMissing); err == nil {
		t.Errorf("Checker did not report error for missing file")
	}
}

