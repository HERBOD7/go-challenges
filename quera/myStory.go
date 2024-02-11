package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func remainsDay(month, day int) int {
	monthDays := 0
	if month <= 6 {
		monthDays = 31
	} else if month > 6 && month <= 11 {
		monthDays = 30
	} else {
		monthDays = 29
	}

	diffDays := monthDays - day

	return diffDays
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p1 := strings.Split(scanner.Text(), " ")

	p1Mon, _ := strconv.Atoi(p1[0])
	p1Day, _ := strconv.Atoi(p1[1])

	scanner.Scan()
	p2 := strings.Split(scanner.Text(), " ")

	p2Mon, _ := strconv.Atoi(p2[0])
	p2Day, _ := strconv.Atoi(p2[1])

	p1RemainDays := remainsDay(p1Mon, p1Day)

	result := 0
	monthDiff := p2Mon - p1Mon
	absDiff := math.Abs(float64(monthDiff))
	if absDiff > 1 {
		daysSum := 0
		for i := 1; i < int(absDiff); i++ {
			daysSum += remainsDay(p1Mon+i, 0)
		}
		result = daysSum + p1RemainDays + p2Day
	} else if absDiff == 1 {
		result = p1RemainDays + p2Day
	} else {
		daysDiff := math.Abs(float64(p2Day - p1Day))
		result = int(daysDiff)
	}

	fmt.Println(result)
}
