package release

import (
	"git-release-materials/argument"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func OutputMaterials(changeList ChangeList, args argument.Args, outputDirPath string) {
	releaseDirPath := createReleaseDir(outputDirPath)
	outputFilePaths := createOutputFilePaths(changeList)

	for _, outputFilePath := range outputFilePaths {
		outputFile(args, releaseDirPath, outputFilePath)
	}
}

func createReleaseDir(outputDirPath string) string {
	dirPath := filepath.Join(outputDirPath, "Release")
	if err := os.MkdirAll(dirPath, 0744); err != nil {
		log.Fatal(err)
	}

	return dirPath
}

func createOutputFilePaths(changeList ChangeList) []string {
	list := append(changeList.Added, changeList.Modified...)
	list = append(list, changeList.Renamed...)
	list = append(list, changeList.Copied...)

	return list
}

func outputFile(args argument.Args, releaseDirPath string, outputFilePath string) {
	target := args.Commit2 + ":" + outputFilePath

	content, err := exec.Command("git", "cat-file", "-p", target).Output()
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filepath.Join(releaseDirPath, outputFilePath), content, 0644); err != nil {
		log.Fatal(err)
	}
}
