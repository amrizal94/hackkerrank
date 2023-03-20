package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(candles []int32) int32 {
	var sum int32
	sort.Slice(candles, func(i, j int) bool { return candles[i] < candles[j] })
	for _, v := range candles {
		if v == candles[len(candles)-1] {
			sum++
		}
	}
	return sum
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}))
}
