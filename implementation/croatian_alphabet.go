// TODO: 크로아티아 알파벳
// https://www.acmicpc.net/problem/2941

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	croatianLetters := []string{"dz=", "lj", "nj", "c=", "c-", "d-", "s=", "z="}
	count := 0

	for i := 0; i < len(input); {
		found := false
		for _, croatian := range croatianLetters {
			if strings.HasPrefix(input[i:], croatian) {
				count++
				i += len(croatian)
				found = true
				break
			}
		}
		if !found {
			count++
			i++
		}
	}
	fmt.Println(count)
}
