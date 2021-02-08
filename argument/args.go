package argument

import (
	"errors"
	"github.com/jessevdk/go-flags"
	"os"
)

type Args struct {
	Command   string
	Commit1   string
	Commit2   string
	GitDir    string `short:"g" long:"git-dir" description:"git root directory"`
	OutputDir string `short:"o" long:"output-dir" description:"output directory"`
}

func Parse() (Args, error) {
	args := Args{}
	nonFlagArgs, err := flags.Parse(&args)
	if err != nil {
		return args, err
	}

	if len(nonFlagArgs) != 3 {
		return args, errors.New("at least three arguments are required, ex. git-release-materials sub-command commit1 commit2")
	}

	args.Command = nonFlagArgs[0]
	args.Commit1 = nonFlagArgs[1]
	args.Commit2 = nonFlagArgs[2]

	if args.GitDir == "" {
		args.GitDir, _ = os.Getwd()
	}
	if args.OutputDir == "" {
		args.OutputDir, _ = os.Getwd()
	}

	return args, nil
}
