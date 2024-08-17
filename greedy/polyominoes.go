// TODO: 폴리오미노
// https://www.acmicpc.net/problem/1343

package main

import (
	"bufio"
	"fmt"
	"os"
)

func coverWithPolyominoes(board string) string {
	n := len(board)
	result := make([]rune, n)

	i := 0
	for i < n {
		if board[i] == '.' {
			result[i] = '.'
			i++
			continue
		}

		countX := 0
		for i+countX < n && board[i+countX] == 'X' {
			countX++
		}

		if countX%2 != 0 {
			return "-1"
		}

		for countX > 0 {
			if countX >= 4 {
				for j := 0; j < 4; j++ {
					result[i+j] = 'A'
				}
				i += 4
				countX -= 4
			} else if countX == 2 {
				for j := 0; j < 2; j++ {
					result[i+j] = 'B'
				}
				i += 2
				countX -= 2
			}
		}
	}
	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	board := scanner.Text()

	result := coverWithPolyominoes(board)
	fmt.Println(result)
}
