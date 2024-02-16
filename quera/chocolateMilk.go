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
	childsCap := strings.Split(scanner.Text(), " ")

	result := 0
	for _, capacity := range childsCap {
		childCap, _ := strconv.Atoi(capacity)
		result += childCap
	}

	if result <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(result)
	}
}
