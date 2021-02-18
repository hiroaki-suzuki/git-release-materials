package main

import (
	"errors"
	"os"

	"github.com/jessevdk/go-flags"
)

type Args struct {
	Command   string
	Commit1   string
	Commit2   string
	GitDir    string `short:"g" long:"git-dir" description:"git root directory"`
	OutputDir string `short:"o" long:"output-dir" description:"output directory"`
	Exclude   string `short:"e" long:"exclude" description:"exclude the specified files"`
}

func Parse() (Args, error) {
	args := Args{}
	nonFlagArgs, err := flags.Parse(&args)
	if err != nil {
		return args, err
	}

	if len(nonFlagArgs) < 2 {
		return args, errors.New("at least two arguments are required, ex. git-release-materials sub-command commit1")
	}

	args.Command = nonFlagArgs[0]
	args.Commit1 = nonFlagArgs[1]
	if len(nonFlagArgs) >= 3 {
		args.Commit2 = nonFlagArgs[2]
	} else {
		args.Commit2 = "HEAD"
	}

	if args.GitDir == "" {
		args.GitDir, _ = os.Getwd()
	}
	if args.OutputDir == "" {
		args.OutputDir, _ = os.Getwd()
	}

	return args, nil
}
