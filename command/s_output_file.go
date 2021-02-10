package command

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func createOutputDir(outputDirPath string, dirName string) string {
	dirPath := filepath.Join(outputDirPath, dirName)
	if err := os.MkdirAll(dirPath, 0744); err != nil {
		log.Fatal(err)
	}

	return dirPath
}

func createOutputFilePaths(filePathsList ...[]string) []string {
	var ret []string
	for _, filePaths := range filePathsList {
		ret = append(ret, filePaths...)
	}

	return ret
}

func outputFile(commit string, releaseDirPath string, outputFilePath string) {
	target := commit + ":" + outputFilePath

	content, err := exec.Command("git", "cat-file", "-p", target).Output()
	if err != nil {
		log.Fatal("fail git cat-file. ", err)
	}

	path := filepath.Join(releaseDirPath, outputFilePath)
	parentDirPath := filepath.Dir(path)
	if _, err := os.Stat(parentDirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDirPath, 0744); err != nil {
			log.Fatal(err)
		}
	}

	if err := ioutil.WriteFile(filepath.Join(releaseDirPath, outputFilePath), content, 0744); err != nil {
		log.Fatal(err)
	}
}
