package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, c := range S {
		if strings.ContainsRune(J, c) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aaaa"))
}
