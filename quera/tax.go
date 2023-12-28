package main

import (
	"fmt"
	"math"
)

func main() {
	var input, step int
	var tax float64
	fmt.Scanf("%d", &input)
	for int(input) > 0 {
		switch step {
		case 1:
			if input > 10000 {
				tax += 500
				input -= 10000
			} else {
				tax += math.Floor(float64(input) * 0.05)
				input = 0
			}
			step += 1
		case 2:
			if input > 40000 {
				tax += 4000
				input -= 40000
			} else {
				tax += float64(input) * 0.1
				input = 0
			}
			step += 1
		case 3:
			if input > 50000 {
				tax += 7500
				input -= 50000
			} else {
				tax += float64(input) * 0.15
				input = 0
			}
			step += 1
		case 4:
			if input > 100000 {
				tax += 20000
				input -= 100000
			} else {
				tax += float64(input) * 0.2
				input = 0
			}
		default:
			step += 1
		}
	}
	fmt.Println(int(tax))
}
