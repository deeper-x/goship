package conf

import (
	"fmt"
	"os"
)

// ProjectRoot base directory
var ProjectRoot = fmt.Sprintf("%s/src/github.com/deeper-x/goship", os.Getenv("GOPATH"))

// EnvFile .env location
var EnvFile = fmt.Sprintf("%s/.env", ProjectRoot)

// PRegisterSQL register dot file
var PRegisterSQL = fmt.Sprintf("%s/qsql/register.sql", ProjectRoot)

// PLiveSQL live data dot file
var PLiveSQL = fmt.Sprintf("%s/qsql/realtime.sql", ProjectRoot)

// PMeteoSQL meteo data dot file
var PMeteoSQL = fmt.Sprintf("%s/qsql/meteo.sql", ProjectRoot)
