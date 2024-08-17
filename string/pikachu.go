// TODO: 피카츄
// https://www.acmicpc.net/problem/14405

package main

import (
	"fmt"
	"strings"
)

func isPikachuLanguage(s string) bool {
	for len(s) > 0 {
		if strings.HasPrefix(s, "pi") {
			s = s[2:]
		} else if strings.HasPrefix(s, "ka") {
			s = s[2:]
		} else if strings.HasPrefix(s, "chu") {
			s = s[3:]
		} else {
			return false
		}
	}
	return true
}

func main() {
	var S string
	fmt.Scanf("%s", &S)

	if isPikachuLanguage(S) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
