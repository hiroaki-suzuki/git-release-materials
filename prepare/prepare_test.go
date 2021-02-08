package prepare

import (
	"git-release-materials/argument"
	"testing"
)

func TestPrepare(t *testing.T) {
	args := argument.Args{GitDir: "../"}
	err := Prepare(args)

	if err != nil {
		t.Errorf("%v", err)
	}
}
