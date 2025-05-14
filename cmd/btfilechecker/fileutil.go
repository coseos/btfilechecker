package main

import "os"

// IsValidFolder checks if the given path is a folder.
func IsValidFolder(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}

// FileExists checks if a file exists at the given path.
func FileExists(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return !info.IsDir()
}

