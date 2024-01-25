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
	input := scanner.Text()
	degrees := strings.Split(input, " ")
	var sum int
	var result string
	for index, value := range degrees {
		valueInt, _ := strconv.Atoi(value)
		if valueInt <= 0 {
			result = "No"
		}
		sum += valueInt

		if index+1 == len(degrees) {
			if sum > 180 || sum < 180 {
				result = "No"
			}
		}
	}

	if result != "No" {
		result = "Yes"
	}

	fmt.Println(result)
}
