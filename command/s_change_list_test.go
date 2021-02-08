package command

import "testing"

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
