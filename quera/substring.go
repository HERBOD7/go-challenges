package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string
	firstSub := strings.Split("HAFEZ", "")
	secondSub := strings.Split("MOLANA", "")

	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	for i, line := range inputs {
		f := 0
		s := 0
		//lineLen := len(line)
		for _, char := range line {
			charStr := string(char)
			if f < len(firstSub) && charStr == firstSub[f] {
				f++
			} else if s < len(secondSub) && charStr == secondSub[s] {
				s++
			}
		}

		if f == len(firstSub) || s == len(secondSub) {
			fmt.Printf("%d ", i+1)
		}
	}
}
