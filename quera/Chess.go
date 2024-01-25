package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	chessPieces := []int{1, 1, 2, 2, 2, 8}
	var result string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	input := strings.Split(line, " ")

	for index, item := range input {
		chessPiece, _ := strconv.Atoi(item)
		diff := chessPieces[index] - chessPiece
		result += fmt.Sprintf("%d ", diff)
	}

	fmt.Println(strings.Trim(result, " "))
}
