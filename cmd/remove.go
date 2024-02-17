package cmd

import (
	"cyrusn/rummikub/db"
	"flag"
	"fmt"
	"os"
)

var removeCmd = flag.NewFlagSet("remove", flag.ExitOnError)

func Remove() {
	var round = removeCmd.Int("r", 0, "Remove all scores in give round")
	var filePath = removeCmd.String("f", DEFAULT_FILE_PATH, FILE_PATH_HELP)
	removeCmd.Parse(os.Args[2:])

	sqlDb := Connect(*filePath, false)

	if *round == 0 {
		r, err := db.GetLastRound(sqlDb)
		handleError(err)
		round = &r
	}
	err := db.RemoveRoundScores(sqlDb, *round)
	handleError(err)

	fmt.Println("Round", *round, "removed")
}
