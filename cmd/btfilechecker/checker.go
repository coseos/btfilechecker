package main

import (
    "fmt"
    "path/filepath"
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
    fullPath := filepath.Join(c.basePath, entry.Filename)
    if !FileExists(fullPath) {
        return fmt.Errorf("file not found: %s", fullPath)
    }
    return nil
}

