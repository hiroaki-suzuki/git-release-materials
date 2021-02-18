package main

import (
	"log"
	"time"
)

func main() {
	args, err := Parse()
	if err != nil {
		log.Fatal(err)
	}

	subcommand, err := GetSubcommand(args)
	if err != nil {
		log.Fatal(err)
	}

	if err := Prepare(args); err != nil {
		log.Fatal(err)
	}

	outputDirPath, err := CreateOutputRootDir(args, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	changeList := NewChangeList(args)
	subcommand(changeList, outputDirPath)
}
