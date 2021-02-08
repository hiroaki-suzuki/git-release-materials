package command

import (
	"bufio"
	"fmt"
	"git-release-materials/argument"
	"log"
	"os"
	"path/filepath"
)

func OutputList(changeList ChangeList, args argument.Args, outputDirPath string) {
	filePath := filepath.Join(outputDirPath, "changelist.md")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writeString(writer, []string{"# Change List"})
	writeString(writer, []string{fmt.Sprintf("List of changes from %s to %s.", args.Commit1, args.Commit2)})

	writeList(writer, "## Added list", changeList.Added)
	writeList(writer, "## Modified list", changeList.Modified)
	writeList(writer, "## Deleted list", changeList.Deleted)
	writeList(writer, "## Renamed list", changeList.Renamed)
	writeList(writer, "## Copied list", changeList.Copied)
	writeList(writer, "## Broken list", changeList.Broken)

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
