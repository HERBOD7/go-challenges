//input
//5 4
//3 4 1 2 2

//output
//2

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
	jarAndCap := strings.Split(scanner.Text(), " ")

	scanner.Scan()
	jams := strings.Split(scanner.Text(), " ")

	jarCount, _ := strconv.Atoi(jarAndCap[0])
	capacity, _ := strconv.Atoi(jarAndCap[1])

	totalCap := jarCount * capacity
	totalJam := 0

	for _, value := range jams {
		intValue, _ := strconv.Atoi(value)
		totalJam += intValue
	}

	remainCap := totalCap - totalJam
	remainJar := remainCap / capacity
	fmt.Println(remainJar)
}
