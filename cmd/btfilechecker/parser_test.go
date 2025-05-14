package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBTFileParser_Parse(t *testing.T) {
	// Prepare a temporary file with test data
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_input.txt")
	content := `
# This is a comment
BT: "First" file1.txt
BT: "Second" file2.txt

# Another comment
Invalid line
BT: "Third" file3.txt
`
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write test input file: %v", err)
	}

	parser := NewBTFileParser(testFile)
	entries, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned error: %v", err)
	}

	// We expect 3 valid entries
	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	tests := []struct {
		name     string
		filename string
	}{
		{"First", "file1.txt"},
		{"Second", "file2.txt"},
		{"Third", "file3.txt"},
	}

	for i, want := range tests {
		got := entries[i]
		if got.Name != want.name {
			t.Errorf("Entry %d: expected Name %q, got %q", i, want.name, got.Name)
		}
		if got.Filename != want.filename {
			t.Errorf("Entry %d: expected Filename %q, got %q", i, want.filename, got.Filename)
		}
		if got.Label != LabelBT {
			t.Errorf("Entry %d: expected Label %q, got %q", i, LabelBT, got.Label)
		}
	}
}

