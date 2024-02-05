package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	switch input {
	case "space":
		fmt.Println("blue")
	case "mind":
		fmt.Println("yellow")
	case "reality":
		fmt.Println("red")
	case "power":
		fmt.Println("purple")
	case "time":
		fmt.Println("green")
	case "soul":
		fmt.Println("orange")
	}
}
