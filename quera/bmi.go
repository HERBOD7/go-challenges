package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	weight, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 64)

	bmi := weight / math.Pow(height, 2)
	fmt.Printf("%.2f", bmi)
	fmt.Println()

	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("Normal")
	} else if bmi >= 25 && bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}
}
