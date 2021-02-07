package main

import (
	"errors"
	"git-release-materials/argument"
	"git-release-materials/release"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	args := getArgs()

	changeDirectory(args)
	verifyGitRoot(args)

	outputDirPath := createOutputDir(args)
	changeList := release.NewChangeList(args)

	switch args.Command {
	case "list":
		release.OutputList(changeList, args, outputDirPath)
	case "release":
		release.OutputList(changeList, args, outputDirPath)
		release.OutputMaterials(changeList, args, outputDirPath)
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

func verifyGitRoot(args argument.Args) {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Fatal("the specified directory is not git root. ", args.WorkDir)
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
