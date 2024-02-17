package cmd

import (
	"cyrusn/rummikub/db"
	"flag"
	"fmt"
	"os"
)

type Result struct {
	Username string
	Round    int
	Score    int
}

var resultCmd = flag.NewFlagSet("result", flag.ExitOnError)

func GetResult() {
	var filePath = resultCmd.String("f", DEFAULT_FILE_PATH, FILE_PATH_HELP)
	resultCmd.Parse(os.Args[2:])

	sqlDb := Connect(*filePath, false)

	lastRound, _ := db.GetLastRound(sqlDb)
	results, err := db.GetResults(sqlDb)
	handleError(err)

	fmt.Println("Round:", lastRound)
	fmt.Println("----------")

	for _, result := range results {
		fmt.Println(result.Username, result.Mark)
	}

}
