package conf

import (
	"regexp"
	"testing"
)

func TestParameters(t *testing.T) {
	matched, _ := regexp.MatchString(`.*/src/github.com/deeper-x/goship`, ProjectRoot)

	if !matched {
		t.Error("ProjectRoot")
	}

	matched, _ = regexp.MatchString(`.*/\.env`, EnvFile)

	if !matched {
		t.Error("EnvFile")
	}

	matched, _ = regexp.MatchString(`.*/qsql/register.sql`, PRegisterSQL)

	if !matched {
		t.Error("ProjectRoot")
	}

	matched, _ = regexp.MatchString(`.*/qsql/realtime.sql`, PLiveSQL)

	if !matched {
		t.Error("ProjectRoot")
	}

	matched, _ = regexp.MatchString(`.*/qsql/meteo.sql`, PMeteoSQL)

	if !matched {
		t.Error("ProjectRoot")
	}

}
