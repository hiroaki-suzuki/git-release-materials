package command

import (
	"bytes"
	"context"
	"git-release-materials/argument"
	"github.com/codeclysm/extract"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func output(commit1 string, commit2 string, outputDirPath string) {
	diffList, err := createGitDiffList(commit1, commit2)
	if err != nil {
		log.Fatal(err)
	}

	archiveFilePath := filepath.Join(outputDirPath, "archive.zip")
	ret, err := execGitArchive(commit2, diffList, archiveFilePath)
	if err != nil {
		log.Fatal(err, ret)
	}

	data, _ := ioutil.ReadFile(archiveFilePath)
	if err = extract.Zip(context.Background(), bytes.NewBuffer(data), outputDirPath, nil); err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(archiveFilePath); err != nil {
		log.Fatal(err)
	}
}
