package main

import "fmt"

func countApplesAndOranges(s, t, a, b int32, apples, oranges []int32) {
	appleCount, orangeCount := 0, 0
	for i := 0; i < len(apples) || i < len(oranges); i++ {
		if i < len(apples) && a+apples[i] >= s && a+apples[i] <= t {
			appleCount++
		}
		if i < len(oranges) && b+oranges[i] >= s && b+oranges[i] <= t {
			orangeCount++
		}
	}
	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

func main() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
