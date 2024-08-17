// TODO: 회의실 배정
// https://www.acmicpc.net/problem/1931

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Meeting struct {
	start, end int
}

// 회의를 정렬하는 함수
func sortMeetings(meetings []Meeting) {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].end == meetings[j].end {
			return meetings[i].start < meetings[j].start
		}
		return meetings[i].end < meetings[j].end
	})
}

// 최대 회의 수를 계산하는 함수
func maxMeetingCount(meetings []Meeting) int {
	count := 0
	lastEndTime := 0

	for _, meeting := range meetings {
		if meeting.start >= lastEndTime {
			lastEndTime = meeting.end
			count++
		}
	}

	return count
}

func readInput(scanner *bufio.Scanner) []Meeting {
	scanner.Scan() // 첫 번째 줄 읽기 (회의 수)
	N, _ := strconv.Atoi(scanner.Text())

	meetings := make([]Meeting, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		var start, end int
		fmt.Sscanf(line, "%d %d", &start, &end)
		meetings[i] = Meeting{start, end}
	}
	return meetings
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	meetings := readInput(scanner)
	sortMeetings(meetings)
	result := maxMeetingCount(meetings)

	fmt.Println(result)
}
