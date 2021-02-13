package prepare

import (
	"git-release-materials/argument"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPrepare(t *testing.T) {
	args := argument.Args{GitDir: "../"}

	actual := Prepare(args)

	if actual != nil {
		t.Errorf("Prepare(%+v): got %v, want nil", args, actual)
	}
}

func TestPrepareNotGitDir(t *testing.T) {
	args := argument.Args{GitDir: "../../"}

	actual := Prepare(args)

	if actual == nil {
		t.Errorf("Prepare(%+v): got %v, want not nil", args, actual)
	}
}

func TestCreateOutputDir(t *testing.T) {
	testCases := []struct {
		outputDir     string
		now           time.Time
		outputDirName string
	}{
		{"./", time.Date(2006, 1, 2, 3, 4, 5, 1, time.UTC), "grm_20060102_030405"},
		{"./", time.Date(2021, 10, 11, 12, 13, 14, 15, time.UTC), "grm_20211011_121314"},
	}

	for _, testCase := range testCases {
		args := argument.Args{OutputDir: testCase.outputDir}
		now := testCase.now

		currentDir, _ := os.Getwd()
		expected := filepath.Join(currentDir, testCase.outputDirName)
		actual, _ := CreateOutputDir(args, now)

		if actual != expected {
			t.Errorf("CreateOutputDir(%+vA, %s): got %v, want %v", args, now, actual, expected)
		}
		_ = os.Remove(actual)
	}
}
