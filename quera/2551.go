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
	firstInput, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	secondInput := scanner.Text()

	scanner.Scan()
	thirdInput, _ := strconv.Atoi(scanner.Text())

	switch secondInput {
	case "+":
		fmt.Printf("%d\n", firstInput+thirdInput)
	case "*":
		fmt.Printf("%d\n", firstInput*thirdInput)
	}
}
