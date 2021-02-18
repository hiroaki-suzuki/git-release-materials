package main

import (
	"log"
)

func OutputMaterials(args Args, outputDirPath string) {
	diffList, err := createGitDiffList(args.Commit1, args.Commit2)
	if err != nil {
		log.Fatal(err)
	}

	releaseDirPath, err := CreateOutputDir(outputDirPath, "Release")
	if err != nil {
		log.Fatal(err)
	}

	ret, err := execGitArchiveWithExtract(args.Commit2, diffList, releaseDirPath)
	if err != nil {
		log.Fatal(err, string(ret))
	}
}
