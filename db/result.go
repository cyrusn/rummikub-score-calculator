package db

import (
	"database/sql"
)

type Result struct {
	Username string
	Mark     int
}

func GetResults(db *sql.DB) ([]Result, error) {
	rows, err := db.Query(`
  SELECT 
    username, SUM(score) from scores 
  JOIN players ON 
    players.id=player_id
  GROUP BY
    player_id
  ORDER BY
    round;`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var results []Result

	for rows.Next() {
		var result Result
		err = rows.Scan(&result.Username, &result.Mark)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
