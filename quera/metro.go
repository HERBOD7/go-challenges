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
	lines := []string{}
	eyeToEyeCount := 0
	for i := 0; i < 2; i++ {
		scanner.Scan()
		line := scanner.Text()
		lines = append(lines, line)
	}

	firstLine := strings.Split(lines[0], " ")
	secondLine := strings.Split(lines[1], " ")

	for i := 0; i < len(firstLine); i++ {
		firstLineStatus, _ := strconv.Atoi(firstLine[i])
		secondLineStatus, _ := strconv.Atoi(secondLine[i])
		if firstLineStatus == 1 && secondLineStatus == firstLineStatus {
			eyeToEyeCount++
		}
	}

	fmt.Printf("%d", eyeToEyeCount)
}
