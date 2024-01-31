package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()

	streets := []string{}

	for i := 0; i < 2; i++ {
		scanner.Scan()
		street := scanner.Text()
		streets = append(streets, street)
	}

	straightStreets := strings.Split(streets[0], " ")
	outStreets := 0
	inStreets := 0

	for _, direction := range straightStreets {
		if direction == "1" {
			outStreets++
		} else if direction == "0" {
			inStreets++
		}
	}

	if len(straightStreets) > 1 && outStreets >= 1 && inStreets >= 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
