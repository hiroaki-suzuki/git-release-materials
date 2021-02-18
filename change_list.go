package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/ryanuber/go-glob"
)

type ChangeList struct {
	Added    []string
	Modified []string
	Deleted  []string
	Renamed  []string
	Broken   []string
}

func NewChangeList(args Args) ChangeList {
	return ChangeList{
		Added:    createList(args, "A"),
		Modified: createList(args, "M"),
		Deleted:  createList(args, "D"),
		Renamed:  createList(args, "R"),
		Broken:   createList(args, "B"),
	}
}

func createList(args Args, diffFilter string) []string {
	ret, err := exec.Command("git", "diff", "-z", "--name-only", "--diff-filter="+diffFilter, args.Commit1, args.Commit2).Output()
	if err != nil {
		log.Fatal(err)
	}

	return createTargetList(args, ret)
}

func createTargetList(args Args, gitDiffResult []byte) []string {
	list := strings.FieldsFunc(string(gitDiffResult), func(c rune) bool {
		return c == 0
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
