package command

import (
	"git-release-materials/argument"
	"git-release-materials/prepare"
	"log"
)

func OutputBeforeAfter(args argument.Args, outputDirPath string) {
	beforeAfterDirPath, err := prepare.CreateOutputDir(outputDirPath, "BeforeAfter")
	if err != nil {
		log.Fatal(err)
	}

	outputBefore(args, beforeAfterDirPath)
	outputAfter(args, beforeAfterDirPath)
}

func outputBefore(args argument.Args, beforeAfterDirPath string) {
	beforeDirPath, err := prepare.CreateOutputDir(beforeAfterDirPath, "before_"+args.Commit1)
	if err != nil {
		log.Fatal(err)
	}

	output(args.Commit2, args.Commit1, beforeDirPath)
}

func outputAfter(args argument.Args, beforeAfterDirPath string) {
	afterDirPath, err := prepare.CreateOutputDir(beforeAfterDirPath, "after_"+args.Commit2)
	if err != nil {
		log.Fatal(err)
	}

	output(args.Commit1, args.Commit2, afterDirPath)
}

func output(commit1 string, commit2 string, outputDirPath string) {
	diffList, err := createGitDiffList(commit1, commit2)
	if err != nil {
		log.Fatal(err)
	}

	ret, err := execGitArchiveWithExtract(commit2, diffList, outputDirPath)
	if err != nil {
		log.Fatal(err, string(ret))
	}
}
