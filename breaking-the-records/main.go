package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	var minBreaks, maxBreaks int32
	minScore, maxScore := scores[0], scores[0]
	for _, v := range scores {
		if minScore > v {
			minScore = v
			minBreaks++
		} else if maxScore < v {
			maxScore = v
			maxBreaks++
		}
	}
	return []int32{maxBreaks, minBreaks}
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}
