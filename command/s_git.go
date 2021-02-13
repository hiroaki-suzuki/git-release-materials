package command

import (
	"bytes"
	"context"
	"github.com/codeclysm/extract"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func createGitDiffList(commit1 string, commit2 string) ([]string, error) {
	ret, err := exec.Command("git", "diff", "-z", "--name-only", "--diff-filter=d", commit1, commit2).Output()
	if err != nil {
		return nil, err
	}

	return strings.FieldsFunc(string(ret), func(c rune) bool {
		return c == 0
	}), nil
}

func execGitArchiveWithExtract(commit2 string, diffList []string, outputDirPath string) ([]byte, error) {
	archiveFilePath := filepath.Join(outputDirPath, "archive.zip")
	args := []string{"archive", commit2}
	args = append(args, diffList...)
	args = append(args, "--format", "zip", "-o", archiveFilePath)

	ret, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return ret, err
	}

	data, _ := ioutil.ReadFile(archiveFilePath)
	if err = extract.Zip(context.Background(), bytes.NewBuffer(data), outputDirPath, nil); err != nil {
		return []byte{}, err
	}

	err = os.Remove(archiveFilePath)
	return []byte{}, err
}
