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
	if inputInt%2 == 0 {
		fmt.Println("Bala Barare")
	} else {
		fmt.Println("Payin Barare")
	}
}
