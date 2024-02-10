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
	scanner.Text()

	scanner.Scan()
	watermelons := strings.Split(scanner.Text(), " ")

	pointer, _ := strconv.Atoi(watermelons[0])
	result := 0
	for index, watermelon := range watermelons {
		weight, _ := strconv.Atoi(watermelon)
		if weight > pointer {
			pointer = weight
			result = index
		}
	}

	fmt.Println(result + 1)
}
