package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), "")

	fmt.Printf("saal:%s%s\n", input[0], input[1])
	fmt.Printf("maah:%s%s", input[2], input[3])
}
