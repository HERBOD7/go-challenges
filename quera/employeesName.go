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
	namesCount, _ := strconv.Atoi(scanner.Text())

	employees := map[string]int{}

	for i := 0; i < namesCount; i++ {
		scanner.Scan()
		name := strings.Split(scanner.Text(), " ")
		employees[name[0]]++
	}

	result := 0

	for _, value := range employees {
		if value > result {
			result = value
		}
	}

	fmt.Println(result)
}
