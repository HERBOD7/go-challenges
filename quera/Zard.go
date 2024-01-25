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
	var oCount string
	for i := 0; i < inputInt; i++ {
		oCount = oCount + "o"
		_ = i
	}
	fmt.Printf("W%sw!", oCount)
}
