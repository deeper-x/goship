package conf

import (
	"os"
	"testing"
)

func fileExists(ifile string) bool {
	info, err := os.Stat(ifile)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func isDir(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

func TestParameters(t *testing.T) {

	gopath := os.Getenv("GOPATH")

	if len(gopath) == 0 {
		t.Error("gopath not set")
	}

	pr := isDir(ProjectRoot)

	if !pr {
		t.Errorf("%s doesn't exists", ProjectRoot)
	}

	ef := fileExists(EnvFile)

	if !ef {
		t.Errorf("%s doesn't exists", EnvFile)
	}

	prs := fileExists(PRegisterSQL)

	if !prs {
		t.Errorf("%s doesn't exists", PRegisterSQL)
	}

	pls := fileExists(PLiveSQL)

	if !pls {
		t.Errorf("%s doesn't exists", PLiveSQL)
	}

	pms := fileExists(PMeteoSQL)

	if !pms {
		t.Errorf("%s doesn't exists", PMeteoSQL)
	}

}
