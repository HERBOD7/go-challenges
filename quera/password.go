package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gearsCount, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	pass := strings.Split(scanner.Text(), "")

	turns := 0

	for i := 0; i < gearsCount; i++ {
		scanner.Scan()
		gear := strings.Split(scanner.Text(), "")
		for index, item := range gear {
			number, _ := strconv.Atoi(item)
			passNumber, _ := strconv.Atoi(pass[i])

			if number == passNumber {
				if index <= 5 {
					turns += index
				} else {
					turns += 9 - index
				}
				break
			}
		}
	}

	fmt.Println(turns)
}
