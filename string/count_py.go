package main

import (
	"fmt"
	"strings"
)

func countPY(s string) bool {
	// 소문자 변환
	s = strings.ToLower(s)

	pCount := strings.Count(s, "p")
	yCount := strings.Count(s, "y")

	if pCount == yCount {
		return true
	}
	return false
}

func main() {
	fmt.Println(countPY("pPoooyY")) // true
	fmt.Println(countPY("Pyy"))     // false
}
