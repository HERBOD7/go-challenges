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
	input := scanner.Text()
	inputInt, _ := strconv.Atoi(input)

	for i := 1; i <= inputInt; i++ {
		for j := 1; j <= inputInt; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}
