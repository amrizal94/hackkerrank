package main

import "fmt"

func kangoroo(x1, v1, x2, v2 int32) string {

	if v1-v2 == 0 {
		return "NO"
	}
	if (x2-x1)%(v1-v2) == 0 && v1-v2 > 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(kangoroo(43, 2, 70, 2))
}
