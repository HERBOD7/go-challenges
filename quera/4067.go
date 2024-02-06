// https://quera.org/problemset/4067
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
	questions, _ := strconv.Atoi(scanner.Text())

	locations := []string{}
	for i := 0; i < questions; i++ {
		scanner.Scan()
		locations = append(locations, scanner.Text())
	}

	for _, item := range locations {
		location := strings.Split(item, " ")
		x, _ := strconv.Atoi(location[0])
		y, _ := strconv.Atoi(location[1])

		// TODO: fix logic
		if x%2 == 0 && y%2 == 0 {
			fmt.Printf("%d\n", x+y)
		} else if x%2 != 0 && y%2 != 0 {
			fmt.Printf("%d\n", x+y-1)
		} else {
			fmt.Printf("%d\n", -1)
		}
	}
}
