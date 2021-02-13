package command

import (
	"bytes"
	"context"
	"git-release-materials/argument"
	"github.com/codeclysm/extract"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func OutputBeforeAfter(args argument.Args, outputDirPath string) {
	beforeAfterDirPath := createOutputDir(outputDirPath, "BeforeAfter")
	outputBefore(args, beforeAfterDirPath)
	outputAfter(args, beforeAfterDirPath)
}

func outputBefore(args argument.Args, beforeAfterDirPath string) {
	beforeDirPath := createOutputDir(beforeAfterDirPath, "before_"+args.Commit1)
	output(args.Commit2, args.Commit1, beforeDirPath)
}

func outputAfter(args argument.Args, beforeAfterDirPath string) {
	afterDirPath := createOutputDir(beforeAfterDirPath, "after_"+args.Commit2)
	output(args.Commit1, args.Commit2, afterDirPath)
}

func output(commit1 string, commit2 string, afterDirPath string) {
	diffRet, err := exec.Command("git", "diff", "-z", "--name-only", "--diff-filter=d", commit1, commit2).Output()
	if err != nil {
		log.Fatal(err)
	}

	list := strings.FieldsFunc(string(diffRet), func(c rune) bool {
		return c == 0
	})

	archiveFile := afterDirPath + "/archive.zip"
	args := []string{"archive", commit2}
	args = append(args, list...)
	args = append(args, "-o", archiveFile)
	if err := exec.Command("git", args...).Run(); err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadFile(archiveFile)
	buffer := bytes.NewBuffer(data)
	err = extract.Zip(context.Background(), buffer, afterDirPath+"/", nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(archiveFile); err != nil {
		log.Fatal(err)
	}
}
