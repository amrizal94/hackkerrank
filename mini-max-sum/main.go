package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	var sum int
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	for _, v := range arr {
		sum += int(v)
	}
	fmt.Println(sum-int(arr[len(arr)-1]), sum-int(arr[0]))
}
func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})

}
