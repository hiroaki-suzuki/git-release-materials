package main

import (
	"errors"
	"git-release-materials/argument"
	"git-release-materials/command"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	args := getArgs()

	executionFunc, err := getExecutionFunc(args)
	if err != nil {
		log.Fatal(err)
	}

	outputDirPath := createOutputDir(args)

	changeDirectory(args)
	verifyGitRoot(args)

	changeList := command.NewChangeList(args)
	executionFunc(changeList, outputDirPath)
}

func getArgs() argument.Args {
	args, err := argument.GetArgs()
	if err != nil {
		log.Fatal(err)
	}

	return args
}

func getExecutionFunc(args argument.Args) (func(changeList command.ChangeList, outputDirPath string), error) {
	switch args.Command {
	case "list":
		return func(changeList command.ChangeList, outputDirPath string) {
			command.OutputList(changeList, args, outputDirPath)
		}, nil
	case "before-after":
		return func(changeList command.ChangeList, outputDirPath string) {
			command.OutputList(changeList, args, outputDirPath)
			command.OutputBeforeAfter(changeList, args, outputDirPath)
		}, nil
	case "release":
		return func(changeList command.ChangeList, outputDirPath string) {
			command.OutputList(changeList, args, outputDirPath)
			command.OutputBeforeAfter(changeList, args, outputDirPath)
			command.OutputMaterials(changeList, args, outputDirPath)
		}, nil
	default:
		return nil, errors.New("the specified subcommand is not supported. " + args.Command)
	}
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
	outputDirName := "grm_" + time.Now().Format("20060102_030405")
	absOutputDirPath, err := filepath.Abs(args.OutputDir)
	if err != nil {
		log.Fatal(err)
	}

	outputDirPath := filepath.Join(absOutputDirPath, outputDirName)
	if err := os.MkdirAll(outputDirPath, 0744); err != nil {
		log.Fatal(err)
	}

	return outputDirPath
}
