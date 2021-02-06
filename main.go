package main

import (
	"errors"
	"git-release-materials/argument"
	"io/ioutil"
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

	switch args.Command {
	case "list":
		list := getList(args)

		outputDir := "grm_" + time.Now().Format("20060102_030405")
		outputDirPath := filepath.Join(args.OutputDir, outputDir)
		err := os.MkdirAll(outputDirPath, 0744)
		if err = ioutil.WriteFile(filepath.Join(outputDirPath, "list.txt"), []byte(list), 0644); err != nil {
			log.Fatal(err)
		}
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
	err := os.Chdir(args.WorkDir)
	if err != nil {
		log.Fatal(err)
	}
}

func verifyManagedByGit(args argument.Args) {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Fatal("the specified directory is not under git control. ", args.WorkDir)
	}
}

func getList(args argument.Args) string {
	out, err := exec.Command("git", "diff", "--name-only", "--diff-filter=d", args.Commit1, args.Commit2).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
