package command

import (
	"git-release-materials/argument"
	"log"
)

func OutputMaterials(args argument.Args, outputDirPath string) {
	diffList, err := createGitDiffList(args.Commit1, args.Commit2)
	if err != nil {
		log.Fatal(err)
	}

	releaseDirPath := createOutputDir(outputDirPath, "Release")
	ret, err := execGitArchiveWithExtract(args.Commit2, diffList, releaseDirPath)
	if err != nil {
		log.Fatal(err, string(ret))
	}
}
