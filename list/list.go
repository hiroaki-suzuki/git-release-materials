package list

import (
	"bufio"
	"fmt"
	"git-release-materials/argument"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	createList := func(args argument.Args, diffFilter string) []string {
		ret, err := exec.Command("git", "diff", "--name-only", "--diff-filter="+diffFilter, args.Commit1, args.Commit2).Output()
		if err != nil {
			log.Fatal(err)
		}

		return strings.FieldsFunc(string(ret), func(c rune) bool {
			return c == '\n'
		})
	}

	changeList := ChangeList{
		Added:    createList(args, "A"),
		Modified: createList(args, "M"),
		Deleted:  createList(args, "D"),
		Renamed:  createList(args, "R"),
		Copied:   createList(args, "C"),
		Broken:   createList(args, "B"),
	}

	return changeList
}

func (cl ChangeList) Output(args argument.Args, outputDirPath string) {
	filePath := filepath.Join(outputDirPath, "changelist.md")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writeString(writer, []string{"# Change List"})
	writeString(writer, []string{fmt.Sprintf("List of changes from %s to %s.", args.Commit1, args.Commit2)})

	writeList(writer, "## Added list", cl.Added)
	writeList(writer, "## Modified list", cl.Modified)
	writeList(writer, "## Deleted list", cl.Deleted)
	writeList(writer, "## Renamed list", cl.Renamed)
	writeList(writer, "## Copied list", cl.Copied)
	writeList(writer, "## Broken list", cl.Broken)

	if err = writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func writeList(writer *bufio.Writer, title string, lines []string) {
	markdownList := make([]string, len(lines))
	for i, line := range lines {
		markdownList[i] = "* " + line
	}

	writeString(writer, []string{title})
	writeString(writer, markdownList)
}

func writeString(writer *bufio.Writer, lines []string) {
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := fmt.Fprintln(writer, ""); err != nil {
		log.Fatal(err)
	}
}
