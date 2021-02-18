package main

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	currentDir, _ := os.Getwd()

	testCases := []struct {
		osArgs   []string
		expected Args
	}{
		{
			osArgs: []string{"cmd", "list", "tag1", "tag2", "-g", "git-dir", "-o", "output-dir", "-e", ".gitignore,*.md"},
			expected: Args{
				Command:   "list",
				Commit1:   "tag1",
				Commit2:   "tag2",
				GitDir:    "git-dir",
				OutputDir: "output-dir",
				Exclude:   ".gitignore,*.md",
			},
		},
		{
			osArgs: []string{"cmd", "list", "tag1", "tag2"},
			expected: Args{
				Command:   "list",
				Commit1:   "tag1",
				Commit2:   "tag2",
				GitDir:    currentDir,
				OutputDir: currentDir,
				Exclude:   "",
			},
		},
		{
			osArgs: []string{"cmd", "list", "tag1"},
			expected: Args{
				Command:   "list",
				Commit1:   "tag1",
				Commit2:   "HEAD",
				GitDir:    currentDir,
				OutputDir: currentDir,
				Exclude:   "",
			},
		},
	}

	for i, tc := range testCases {
		os.Args = tc.osArgs
		actual, _ := Parse()
		if actual != tc.expected {
			t.Errorf("test case %d failded: got %v, want %v", i, actual, tc.expected)
		}
	}
}

func TestGetArgsNoArguments(t *testing.T) {
	os.Args = []string{"cmd"}

	_, actual := Parse()

	if actual == nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "nil", actual)
	}
}
