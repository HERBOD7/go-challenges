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
	inputs := strings.Split(scanner.Text(), " ")
	n, _ := strconv.ParseFloat(inputs[0], 64)
	k, _ := strconv.ParseFloat(inputs[1], 64)
	s, _ := strconv.ParseFloat(inputs[2], 64)

	if n*k <= s {
		fmt.Println("Kafie!")
	} else {
		fmt.Println("Na! yeki bayad bere sabzi bekhare")
	}
}
