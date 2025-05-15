package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

// BTFileChecker checks for file existence.
type BTFileChecker struct {
    basePath string
}

// NewBTFileChecker creates a new checker.
func NewBTFileChecker(basePath string) *BTFileChecker {
    return &BTFileChecker{basePath: basePath}
}

// Check verifies if the file exists in the base path.
func (c *BTFileChecker) Check(entry BTFileEntry) error {
    filename := strings.ReplaceAll(entry.Filename, "\\", string(filepath.Separator))
    fullPath := filepath.Join(c.basePath, filename)
    if !FileExists(fullPath) {
        return fmt.Errorf("file not found: %s", fullPath)
    }
    return nil
}

