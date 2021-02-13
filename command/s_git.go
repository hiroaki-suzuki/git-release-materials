package command

import (
	"os/exec"
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

func execGitArchive(commit2 string, diffList []string, archiveFilePath string) ([]byte, error) {
	args := []string{"archive", commit2}
	args = append(args, diffList...)
	args = append(args, "--format", "zip", "-o", archiveFilePath)

	return exec.Command("git", args...).CombinedOutput()
}
