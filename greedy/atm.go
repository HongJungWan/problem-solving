// TODO: ATM
// https://www.acmicpc.net/problem/11399

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	P_str := strings.Fields(scanner.Text())

	P := make([]int, N)
	for i := 0; i < N; i++ {
		P[i], _ = strconv.Atoi(P_str[i])
	}

	sort.Ints(P)

	totalTime := 0
	waitTime := 0
	for i := 0; i < N; i++ {
		waitTime += P[i]
		totalTime += waitTime
	}

	fmt.Println(totalTime)
}
