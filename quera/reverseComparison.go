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
	for i := 0; i < 2; i++ {
		scanner.Scan()
		input := scanner.Text()
		inputs = append(inputs, input)
	}

	firstInput := strings.Split(inputs[0], "")
	secondInput := strings.Split(inputs[1], "")

	for i := 2; i >= 0; i-- {
		if firstInput[i] > secondInput[i] {
			fmt.Printf("%s < %s", inputs[1], inputs[0])
			break
		} else if secondInput[i] > firstInput[i] {
			fmt.Printf("%s < %s", inputs[0], inputs[1])
			break
		} else {
			if i == 0 {
				fmt.Printf("%s = %s", inputs[0], inputs[1])
			}
			continue
		}
	}
}
