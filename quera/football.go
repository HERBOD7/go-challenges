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
	matches, _ := strconv.Atoi(scanner.Text())

	matchesResult := [][]string{}
	for i := 0; i < matches; i++ {
		scanner.Scan()
		result := strings.Split(scanner.Text(), " ")
		matchesResult = append(matchesResult, result)
	}

	for _, match := range matchesResult {
		a, _ := strconv.Atoi(match[0])
		b, _ := strconv.Atoi(match[1])
		c, _ := strconv.Atoi(match[2])
		d, _ := strconv.Atoi(match[3])

		if a+c > b+d {
			fmt.Println("perspolis")
		} else if b+d > a+c {
			fmt.Println("esteghlal")
		} else {
			if c > b {
				fmt.Println("perspolis")
			} else if b > c {
				fmt.Println("esteghlal")
			} else {
				fmt.Println("penalty")
			}
		}
	}
}
