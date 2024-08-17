// TODO: 회사에 있는 사람
// https://www.acmicpc.net/problem/7785

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func processLogs(logs []string) map[string]bool {
	company := make(map[string]bool)
	for _, log := range logs {
		parts := strings.Split(log, " ")
		name, status := parts[0], parts[1]

		if status == "enter" {
			company[name] = true
		} else if status == "leave" {
			delete(company, name)
		}
	}
	return company
}

func readLogs() []string {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Scan()을 사용한 후 남은 newline character를 읽어서 무시한다.
	scanner.Scan()

	logs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		logs = append(logs, scanner.Text())
	}
	return logs
}

func printCurrentPeople(company map[string]bool) {
	names := make([]string, 0, len(company))
	for name := range company {
		names = append(names, name)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(names)))

	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	company := processLogs(readLogs())
	printCurrentPeople(company)
}
