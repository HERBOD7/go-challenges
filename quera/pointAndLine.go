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
	line := scanner.Text()
	inputs := strings.Split(line, " ")

	x1, _ := strconv.Atoi(inputs[0])
	x2, _ := strconv.Atoi(inputs[2])
	y1, _ := strconv.Atoi(inputs[1])
	y2, _ := strconv.Atoi(inputs[3])

	if x1 == x2 {
		fmt.Println("Vertical")
	} else if y1 == y2 {
		fmt.Println("Horizontal")
	} else {
		fmt.Println("Try again")
	}
}
