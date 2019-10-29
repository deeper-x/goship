package conf

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
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

func TestEnv(t *testing.T) {
	err := godotenv.Load(EnvFile)
	if err != nil {
		log.Fatal(err)
	}

	gopath := os.Getenv("GOPATH")

	if len(gopath) == 0 {
		t.Error("gopath not set")
	}

	dbdsn := os.Getenv("DB_DSN")

	if len(dbdsn) == 0 {
		t.Error("DB_DSN not set")
	}
}

func TestParameters(t *testing.T) {

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
