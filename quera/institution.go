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
	counter := scanner.Text()
	teacherCount, _ := strconv.Atoi(counter)
	teachers := []string{}
	for i := 0; i < teacherCount*2; i++ {
		scanner.Scan()
		input := scanner.Text()
		teachers = append(teachers, input)
	}

	for index, item := range teachers {
		if (index+1)%2 != 0 {
			scores := strings.Split(teachers[index+1], " ")
			scoreCount := len(scores)
			var sum int
			for _, score := range scores {
				value, _ := strconv.Atoi(score)
				sum = sum + value
			}
			avg := sum / scoreCount
			var quality string
			if avg >= 80 {
				quality = "Excellent"
			} else if avg >= 60 {
				quality = "Very Good"
			} else if avg >= 40 {
				quality = "Good"
			} else {
				quality = "Fair"
			}
			fmt.Println(item, quality)
		}
	}
}
