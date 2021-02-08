package command

import (
	"git-release-materials/argument"
)

func OutputBeforeAfter(changeList ChangeList, args argument.Args, outputDirPath string) {
	beforeAfterDirPath := createOutputDir(outputDirPath, "BeforeAfter")
	outputBefore(changeList, args, beforeAfterDirPath)
	outputAfter(changeList, args, beforeAfterDirPath)
}

func outputBefore(changeList ChangeList, args argument.Args, beforeAfterDirPath string) {
	beforeDirPath := createOutputDir(beforeAfterDirPath, "before_"+args.Commit1)
	beforeOutputFilePaths := createOutputFilePaths(changeList.Modified, changeList.Renamed, changeList.Copied)
	for _, outputFilePath := range beforeOutputFilePaths {
		outputFile(args.Commit1, beforeDirPath, outputFilePath)
	}
}

func outputAfter(changeList ChangeList, args argument.Args, beforeAfterDirPath string) {
	afterDirPath := createOutputDir(beforeAfterDirPath, "after_"+args.Commit2)
	afterOutputFilePaths := createOutputFilePaths(changeList.Added, changeList.Modified, changeList.Renamed, changeList.Copied)
	for _, outputFilePath := range afterOutputFilePaths {
		outputFile(args.Commit2, afterDirPath, outputFilePath)
	}
}
