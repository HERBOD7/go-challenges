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
	input := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(input[0])
	//m, _ := strconv.Atoi(input[1])
	//_ = m
	dictionary := make(map[string]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		property := strings.Split(scanner.Text(), " ")
		key := property[0]
		value := property[1]

		dictionary[key] = value
	}

	scanner.Scan()
	sentence := strings.Split(scanner.Text(), " ")

	result := ""
	for i, word := range sentence {
		translate, exist := dictionary[word]
		if exist {
			result += translate + " kachal!"
		} else {
			result += "kachal!"
		}

		if i+1 < len(sentence) {
			result += " "
		}
	}
	fmt.Println(result)
}
