package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

const (
	FILE_PATH_HELP    = "The file path the sqlite3 database"
	DEFAULT_FILE_PATH = "rummikub.db"
)

func Connect(filePath string, isSetup bool) *sql.DB {
	if len(os.Args) < 2 {
		Help()
		os.Exit(0)
	}

	if _, err := os.Stat(filePath); !(err == nil || isSetup) {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Database does not exist, Please run `setup` to build one.")
			os.Exit(0)
		}

		handleError(err)
	}

	sqlDb, err := sql.Open("sqlite3", filePath)
	handleError(err)

	return sqlDb
}
