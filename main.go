package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Wiki search engine.")
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	sorted := isAlienSorted(words, order)
	fmt.Println(sorted)
}

func isAlienSorted(words []string, order string) bool {
	alphaOrderMap := make(map[byte]int)
	for i, c := range order {
		alphaOrderMap[byte(c)] = i
	}

	return sort.SliceIsSorted(words, func(i, j int) bool {
		for index := 0; index < len(words[i]) && index < len(words[j]); index++ {
			if words[i][index] == words[j][index] {
				continue
			}
			return alphaOrderMap[words[i][index]] < alphaOrderMap[words[j][index]]
		}
		return len(words[i]) < len(words[j])
	})
}
