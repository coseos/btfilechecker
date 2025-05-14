package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: btfilechecker <folder-path-with-slash/> <input-file>")
        os.Exit(1)
    }

    folderPath := os.Args[1]
    inputFile := os.Args[2]

    if folderPath[len(folderPath)-1] != '/' {
        fmt.Println("Error: Folder path must end with a slash (/)")
        os.Exit(1)
    }

    if !IsValidFolder(folderPath) {
        fmt.Printf("Error: %s is not a valid folder\n", folderPath)
        os.Exit(1)
    }

    parser := NewBTFileParser(inputFile)
    entries, err := parser.Parse()
    if err != nil {
        fmt.Printf("Error parsing input file: %v\n", err)
        os.Exit(1)
    }

    checker := NewBTFileChecker(folderPath)
    missing := false
    for _, entry := range entries {
        err := checker.Check(entry)
        if err != nil {
            fmt.Printf("Missing: %s (%v)\n", entry.Filename, err)
            missing = true
        } else {
            fmt.Printf("Exists: %s\n", entry.Filename)
        }
    }

    if missing {
        os.Exit(2) // Files missing
    }
    os.Exit(0) // Success
}

