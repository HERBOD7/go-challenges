//input : 3 2 1 3
//output: 2

package main

func birthdayCakeCandles(candles []int32) int32 {
	tallestCandle := candles[0]
	countOfTallest := 1

	for i := 1; i < len(candles); i++ {
		if candles[i] > tallestCandle {
			tallestCandle = candles[i]
			countOfTallest = 1
		} else if candles[i] == tallestCandle {
			countOfTallest++
		}
	}

	return int32(countOfTallest)
}
