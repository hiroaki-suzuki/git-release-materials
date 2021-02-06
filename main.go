package main

import (
	"errors"
	"git-release-materials/argument"
	"git-release-materials/list"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	args := getArgs()

	changeDirectory(args)
	verifyManagedByGit(args)

	outputDirPath := createOutputDir(args)
	changeList := list.NewChangeList(args)

	switch args.Command {
	case "list":
		changeList.Output(args, outputDirPath)
	default:
		log.Fatal(errors.New("the specified subcommand is not supported. " + args.Command))
	}
}

func getArgs() argument.Args {
	args, err := argument.GetArgs()
	if err != nil {
		log.Fatal(err)
	}

	return args
}

func changeDirectory(args argument.Args) {
	if err := os.Chdir(args.WorkDir); err != nil {
		log.Fatal(err)
	}
}

func verifyManagedByGit(args argument.Args) {
	if _, err := exec.Command("git", "status").Output(); err != nil {
		log.Fatal("the specified directory is not under git control. ", args.WorkDir)
	}
}

func createOutputDir(args argument.Args) string {
	outputDir := "grm_" + time.Now().Format("20060102_030405")
	outputDirPath := filepath.Join(args.OutputDir, outputDir)
	if err := os.MkdirAll(outputDirPath, 0744); err != nil {
		log.Fatal(err)
	}

	return outputDirPath
}
