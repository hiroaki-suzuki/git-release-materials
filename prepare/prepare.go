package prepare

import (
	"git-release-materials/argument"
	"os"
	"path/filepath"
	"time"
)

func Prepare(args argument.Args) error {
	if err := os.Chdir(args.GitDir); err != nil {
		return err
	}

	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		return err
	}

	return nil
}

func CreateOutputDir(args argument.Args) (string, error) {
	outputDirName := "grm_" + time.Now().Format("20060102_030405")
	absOutputDirPath, err := filepath.Abs(args.OutputDir)
	if err != nil {
		return "", err
	}

	outputDirPath := filepath.Join(absOutputDirPath, outputDirName)
	if err := os.MkdirAll(outputDirPath, 0744); err != nil {
		return "", err
	}

	return outputDirPath, nil
}
