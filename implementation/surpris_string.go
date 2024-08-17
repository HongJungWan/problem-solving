// TODO: 놀라운 문자열
// https://www.acmicpc.net/problem/1972

package main

import (
	"bufio"
	"fmt"
	"os"
)

func isSurprising(s string) bool {
	for d := 0; d < len(s)-1; d++ {
		pairs := make(map[string]bool)

		for i := 0; i < len(s)-d-1; i++ {
			pair := string(s[i]) + string(s[i+d+1])

			if pairs[pair] {
				return false
			}
			pairs[pair] = true
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "*" {
			break
		}

		if isSurprising(line) {
			fmt.Printf("%s is surprising.\n", line)
		} else {
			fmt.Printf("%s is NOT surprising.\n", line)
		}
	}
}
