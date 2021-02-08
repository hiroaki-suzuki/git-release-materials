package main

import (
	"git-release-materials/argument"
	"git-release-materials/command"
	"git-release-materials/prepare"
	"log"
)

func main() {
	args, err := argument.Parse()
	if err != nil {
		log.Fatal(err)
	}

	subcommand, err := command.GetSubcommand(args)
	if err != nil {
		log.Fatal(err)
	}

	if err := prepare.Prepare(args); err != nil {
		log.Fatal(err)
	}

	outputDirPath, err := prepare.CreateOutputDir(args)
	if err != nil {
		log.Fatal(err)
	}

	changeList := command.NewChangeList(args)
	subcommand(changeList, outputDirPath)
}
