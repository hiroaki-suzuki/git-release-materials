package main

import (
	"fmt"
	"git-release-materials/argument"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	args, err := argument.GetArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(args.WorkDir)
	if err != nil {
		log.Fatal()
	}

	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Fatal("This directory is not managed by Git. ", args.WorkDir)
	}

	list := getList(args)
	fmt.Println(list)

	outputDir := "grm_" + time.Now().Format("20060102_030405")
	outputDirPath := filepath.Join(args.OutputDir, outputDir)
	err = os.MkdirAll(outputDirPath, 0744)
	if err = ioutil.WriteFile(filepath.Join(outputDirPath, "list.txt"), []byte(list), 0644); err != nil {
		log.Fatal(err)
	}
}

func getList(args argument.Args) string {
	out, err := exec.Command("git", "diff", "--name-only", "--diff-filter=d", args.Commit1, args.Commit2).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
