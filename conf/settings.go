package conf

import (
	"fmt"
	"os"
)

// ProjectRoot base directory
var ProjectRoot = fmt.Sprintf("%s/src/github.com/deeper-x/goship", os.Getenv("GOPATH"))

// PRegisterSQL register dot file
var PRegisterSQL = fmt.Sprintf("%s/qsql/register_queries.sql", ProjectRoot)

// PLiveSQL live data dot file
var PLiveSQL = fmt.Sprintf("%s/qsql/live_queries.sql", ProjectRoot)
