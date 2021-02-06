package argument

import (
	"os"
	"testing"
)

func TestGetArgsForList(t *testing.T) {
	os.Args = []string{"cmd", "list", "tag1", "tag2", "-d", "/tmp/args", "-o", "/tmp/output"}

	expected := Args{
		Command:   "list",
		Commit1:   "tag1",
		Commit2:   "tag2",
		WorkDir:   "/tmp/args",
		OutputDir: "/tmp/output",
	}
	actual, _ := GetArgs()

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestGetArgsDefault(t *testing.T) {
	os.Args = []string{"cmd", "list", "tag1", "tag2"}

	currentDir, _ := os.Getwd()
	expected := Args{
		Command:   "list",
		Commit1:   "tag1",
		Commit2:   "tag2",
		WorkDir:   currentDir,
		OutputDir: currentDir,
	}
	actual, _ := GetArgs()

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestGetArgsNoArguments(t *testing.T) {
	os.Args = []string{"cmd"}

	_, actual := GetArgs()

	if actual == nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "nil", actual)
	}
}
