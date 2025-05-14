package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

// BTFileEntry represents a single line entry in the input file.
type BTFileEntry struct {
    Label    string
    Name     string
    Filename string
}

// BTFileParser parses the input file.
type BTFileParser struct {
    filename string
}

// NewBTFileParser creates a new parser.
func NewBTFileParser(filename string) *BTFileParser {
    return &BTFileParser{filename: filename}
}

// Parse reads and parses the file, returning valid entries.
func (p *BTFileParser) Parse() ([]BTFileEntry, error) {
    file, err := os.Open(p.filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var entries []BTFileEntry
    scanner := bufio.NewScanner(file)
    lineNum := 0

    // Regex: use the label constant
    re := regexp.MustCompile(fmt.Sprintf(`^%s\s+"([^"]+)"\s+([^\s]+)$`, regexp.QuoteMeta(LabelBT)))

    for scanner.Scan() {
        lineNum++
        line := strings.TrimSpace(scanner.Text())
        if len(line) == 0 || strings.HasPrefix(line, "#") {
            continue // skip comments and blank lines
        }
        matches := re.FindStringSubmatch(line)
        if matches == nil {
            fmt.Printf("Warning: Invalid format on line %d: %s\n", lineNum, line)
            continue
        }
        entry := BTFileEntry{
            Label:    LabelBT,
            Name:     matches[1],
            Filename: matches[2],
        }
        entries = append(entries, entry)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return entries, nil
}

