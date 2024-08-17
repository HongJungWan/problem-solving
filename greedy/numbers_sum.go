// TODO: 수들의 합
// https://www.acmicpc.net/problem/1789

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	S, _ := strconv.Atoi(scanner.Text())

	var n = 1
	var sum = 0

	// 자연수를 더하면서 S를 넘지 않는 최대 자연수 개수 찾기
	for sum+n <= S {
		sum += n
		n++
	}

	// n-1이 최대 자연수의 개수가 됨
	fmt.Println(n - 1)
}

// 합이 S를 처음으로 초과하지 않는 n을 찾을 때까지 자연수를 더하게 되며, n−1까지의 합이 S 이하인 최대값이 된다.
