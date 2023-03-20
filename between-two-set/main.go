package main

import "fmt"

func getTotalX(a, b []int32) int32 {
	max := a[0]
	min := b[0]

	// Find the maximum value in array a and the minimum value in array b
	for _, num := range a {
		if num > max {
			max = num
		}
	}

	for _, num := range b {
		if num < min {
			min = num
		}
	}

	var count int32

	// Check all integers from max to min
	for i := max; i <= min; i++ {
		factorOfA := true
		factorOfB := true

		// Check if all elements in array a are factors of i
		for _, num := range a {
			if i%num != 0 {
				factorOfA = false
				break
			}
		}

		// Check if i is a factor of all elements in array b
		for _, num := range b {
			if num%i != 0 {
				factorOfB = false
				break
			}
		}

		// Increment the count if i satisfies both conditions
		if factorOfA && factorOfB {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}))
}
