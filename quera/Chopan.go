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
	var inputs [2]string
	for i := 0; i < 2; i++ {
		scanner.Scan()
		inputs[i] = scanner.Text()
	}

	financialDetails := strings.Split(inputs[0], " ")
	w, _ := strconv.ParseFloat(financialDetails[0], 64)
	sh, _ := strconv.ParseFloat(financialDetails[1], 64)
	p, _ := strconv.ParseFloat(inputs[1], 64)
	salary := 365 * w
	penalty := p * 365 * sh
	revenue := salary - penalty
	fmt.Printf("%.2f", revenue)
}
