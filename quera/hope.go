package main

import "fmt"

func main() {
	var p, q int
	_, err := fmt.Scanf("%d %d", &p, &q)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	for i := 1; i <= q; i++ {
		printNumOrHope(i, p)
	}
}

func printNumOrHope(index, div int) {
	if index%div == 0 {
		multiple := index / div
		printHopeRepeatedly(multiple)
	} else {
		fmt.Println(index)
	}
}

func printHopeRepeatedly(n int) {
	for j := 0; j < n; j++ {
		if j == n-1 {
			fmt.Println("Hope")
		} else {
			fmt.Print("Hope ")
		}
	}
}
