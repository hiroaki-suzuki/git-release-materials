package command

import (
	"git-release-materials/argument"
	"reflect"
	"testing"
)

func TestCreateTargetList(t *testing.T) {
	args := argument.Args{Exclude: ".gitignore,*.md"}
	gitDiffResult := []byte(".gitignore\nfunctions.go\npath/to/fileB\nREADME.md")

	expected := []string{"functions.go", "path/to/fileB"}
	actual := createTargetList(args, gitDiffResult)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestCanAddTargetList(t *testing.T) {
	excludeList := []string{".gitignore", "*.md", "test/*"}

	successFiles := []string{"fileA", "path/to/fileB", ".editorconfig"}
	for _, file := range successFiles {
		if !canAddTargetList(excludeList, file) {
			t.Errorf("%s should be true.", file)
		}
	}

	failFiles := []string{".gitignore", "README.md", "test/Aaa/test_test.go", "test/foo_test.go"}
	for _, file := range failFiles {
		if canAddTargetList(excludeList, file) {
			t.Errorf("%s should be false.", file)
		}
	}
}
