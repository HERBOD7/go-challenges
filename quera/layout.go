package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	columnsCount, _ := strconv.Atoi(scanner.Text())

	columns := []int{}
	sum := 0
	for i := 0; i < columnsCount; i++ {
		scanner.Scan()
		column, _ := strconv.Atoi(scanner.Text())
		sum += column
		columns = append(columns, column)
	}

	avg := sum / len(columns)
	mints := 0
	for _, column := range columns {
		if column > avg {
			diff := column - avg
			mints += diff
		}
	}

	fmt.Println(mints)
}
