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
	n, _ := strconv.Atoi(inputs[0])
	k, _ := strconv.Atoi(inputs[1])
	s, _ := strconv.Atoi(inputs[2])

	if n*k == s {
		fmt.Println("Kafie!")
	} else {
		fmt.Println("Na! yeki bayad bere sabzi bekhare")
	}
}
