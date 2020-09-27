package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	counter := make(map[rune]int)
	for _, v := range magazine {
		counter[v]++
	}

	for _, v := range ransomNote {
		counter[v]--
	}

	for _, v := range counter {
		if v < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
