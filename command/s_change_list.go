package command

import (
	"git-release-materials/argument"
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

	return strings.FieldsFunc(string(ret), func(c rune) bool {
		return c == '\n'
	})
}
