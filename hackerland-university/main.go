package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	var res []int32

	for _, v := range grades {
		mod := 5 - (v % 5)
		if mod < 3 {
			v += mod
			if v < 40 {
				v -= mod
			}
		}
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(gradingStudents([]int32{37, 38}))
}
