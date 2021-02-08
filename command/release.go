package command

import (
	"git-release-materials/argument"
)

func OutputMaterials(changeList ChangeList, args argument.Args, outputDirPath string) {
	releaseDirPath := createOutputDir(outputDirPath, "Release")
	outputFilePaths := createOutputFilePaths(changeList.Added, changeList.Modified, changeList.Renamed, changeList.Copied)

	for _, outputFilePath := range outputFilePaths {
		outputFile(args.Commit2, releaseDirPath, outputFilePath)
	}
}
