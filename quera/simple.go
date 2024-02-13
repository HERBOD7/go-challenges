package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	inputs := []float64{}

	for i := 0; i < 4; i++ {
		scanner.Scan()
		input, _ := strconv.ParseFloat(scanner.Text(), 64)
		inputs = append(inputs, input)
	}

	var sum float64
	multiple := 1.0
	minInput := inputs[0]
	maxInput := inputs[0]

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		sum += input
		multiple *= input
		if input < minInput {
			minInput = input
		} else if input > maxInput {
			maxInput = input
		}
	}

	fmt.Printf("Sum : %.6f\n", sum)
	fmt.Printf("Average : %.6f\n", sum/4)
	fmt.Printf("Product : %.6f\n", multiple)
	fmt.Printf("MAX : %.6f\n", maxInput)
	fmt.Printf("MIN : %.6f\n", minInput)
}
