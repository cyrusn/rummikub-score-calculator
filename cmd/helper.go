package cmd

import (
	"fmt"
	"log"
	"strconv"
)

func sum(arr ...string) int {
	sum := 0
	idx := 0
	for {
		if idx > len(arr)-1 {
			break
		}
		if arr[idx] == "" {
			idx += 1
			continue
		}

		if mark, err := strconv.Atoi(arr[idx]); err != nil {
			fmt.Println(err)
			idx += 1
			continue
		} else {
			sum += mark
			idx += 1
		}
	}
	return sum
}

func handleError(err error) {
	if err != nil {
		fmt.Println(`Oops! Something went wrong.
Please try again. If problem persists, run setup command to rebuild database.

Error logs:`)
		log.Fatal(err)
	}
}
