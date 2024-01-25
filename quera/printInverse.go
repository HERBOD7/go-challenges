package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	list := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		input, _ := strconv.Atoi(line)
		if input != 0 {
			list = append(list, input)
		}
	}

	listLen := len(list)
	for i := 1; i <= listLen; i++ {
		fmt.Println(list[listLen-i])
	}
}
