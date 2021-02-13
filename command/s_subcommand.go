package command

import (
	"errors"
	"git-release-materials/argument"
)

func GetSubcommand(args argument.Args) (func(changeList ChangeList, outputDirPath string), error) {
	switch args.Command {
	case "list":
		return func(changeList ChangeList, outputDirPath string) {
			OutputList(changeList, args, outputDirPath)
		}, nil
	case "before-after":
		return func(changeList ChangeList, outputDirPath string) {
			OutputList(changeList, args, outputDirPath)
			OutputBeforeAfter(args, outputDirPath)
		}, nil
	case "release":
		return func(changeList ChangeList, outputDirPath string) {
			OutputList(changeList, args, outputDirPath)
			OutputBeforeAfter(args, outputDirPath)
			OutputMaterials(changeList, args, outputDirPath)
		}, nil
	default:
		return nil, errors.New("the specified subcommand is not supported. " + args.Command)
	}
}
