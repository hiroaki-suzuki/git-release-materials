package command

import (
	"os"
	"reflect"
	"testing"
)

func TestCreateGitDiffList(t *testing.T) {
	commit1 := "58a905e"
	commit2 := "9718b8b"

	expected := []string{".gitignore", "argument/args.go", "argument/args_test.go", "go.mod", "go.sum", "main.go"}
	actual, _ := createGitDiffList(commit1, commit2)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("createGitDiffList(%s, %s): got %s, want %s", commit1, commit2, actual, expected)
	}
}

func TestCreateGitDiffListFail(t *testing.T) {
	commit1 := "58a905e"
	commit2 := "unknown"

	_, actual := createGitDiffList(commit1, commit2)

	if actual == nil {
		t.Errorf("createGitDiffList(%s, %s): got %v, want not nil", commit1, commit2, actual)
	}
}

func TestExecGitArchive(t *testing.T) {
	commit2 := "9718b8b"
	diffList := []string{".gitignore", "argument/args.go", "argument/args_test.go", "go.mod", "go.sum", "main.go"}
	archiveFilePath := "archive.zip"

	_ = os.Chdir("../")
	ret, err := execGitArchive(commit2, diffList, archiveFilePath)
	_ = os.Remove(archiveFilePath)

	if err != nil {
		t.Errorf("execGitArchive(%s, %v, %s): an error has occurred. %v %s", commit2, diffList, archiveFilePath, err, ret)
	}
}

func TestExecGitArchiveFail(t *testing.T) {
	commit2 := "9718b8b"
	diffList := []string{".gitignore", "argument/args.go", "argument/args_test.go", "go.mod", "go.sum", "main.go"}
	archiveFilePath := "archive.zip"

	_ = os.Chdir("../../")
	_, err := execGitArchive(commit2, diffList, archiveFilePath)
	_ = os.Remove(archiveFilePath)

	if err == nil || err.Error() == "" {
		t.Errorf("execGitArchive(%s, %v, %s): no error occurred.", commit2, diffList, archiveFilePath)
	}
}
