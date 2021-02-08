package command

import (
	"git-release-materials/argument"
	"github.com/ryanuber/go-glob"
	"log"
	"os/exec"
	"strings"
)

type ChangeList struct {
	Added    []string
	Modified []string
	Deleted  []string
	Renamed  []string
	Copied   []string
	Broken   []string
}

func NewChangeList(args argument.Args) ChangeList {
	return ChangeList{
		Added:    createList(args, "A"),
		Modified: createList(args, "M"),
		Deleted:  createList(args, "D"),
		Renamed:  createList(args, "R"),
		Copied:   createList(args, "C"),
		Broken:   createList(args, "B"),
	}
}

func createList(args argument.Args, diffFilter string) []string {
	ret, err := exec.Command("git", "diff", "--name-only", "--diff-filter="+diffFilter, args.Commit1, args.Commit2).Output()
	if err != nil {
		log.Fatal(err)
	}

	return createTargetList(args, ret)
}

func createTargetList(args argument.Args, gitDiffResult []byte) []string {
	list := strings.FieldsFunc(string(gitDiffResult), func(c rune) bool {
		return c == '\n'
	})

	var excludeList = strings.Split(args.Exclude, ",")
	var targetList []string
	for _, file := range list {
		if canAddTargetList(excludeList, file) {
			targetList = append(targetList, file)
		}
	}

	return targetList
}

func canAddTargetList(excludeList []string, file string) bool {
	for _, excludeFile := range excludeList {
		if glob.Glob(excludeFile, file) {
			return false
		}
	}
	return true
}
