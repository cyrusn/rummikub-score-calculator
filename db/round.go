package db

import "database/sql"

func GetLastRound(db *sql.DB) (int, error) {
	var round int
	err := db.QueryRow("SELECT MAX(round) FROM scores").Scan(&round)
	if err != nil {
		return 0, err
	}
	return round, nil
}
