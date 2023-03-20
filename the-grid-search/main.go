package main

import "fmt"

func gridSearch(G []string, P []string) string {
	// Write your code here
	rowsG := len(G)
	colsG := len(G[0])
	rowsP := len(P)
	colsP := len(P[0])

	for i := 0; i <= rowsG-rowsP; i++ {
		for j := 0; j <= colsG-colsP; j++ {
			if G[i][j:j+colsP] == P[0] {
				match := true
				for k := 1; k < rowsP; k++ {
					if G[i+k][j:j+colsP] != P[k] {
						match = false
						break
					}
				}
				if match {
					return "YES"
				}
			}
		}
	}

	return "NO"
}

func main() {
	fmt.Println(gridSearch([]string{"7283455864", "6731158619", "8988242643", "3830589324", "2229505813", "5633845374", "6473530293", "7053106601", "0834282956", "4607924137"}, []string{"9505", "3845", "3530"}))
}
