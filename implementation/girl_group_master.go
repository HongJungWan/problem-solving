// TODO: 걸그룹 마스터 준석이
// https://www.acmicpc.net/problem/16165

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type groupInfo struct {
	groupName string
	members   []string
}

// 새로운 groupInfo를 생성하는 함수
func newGroupInfo(groupName string, members []string) groupInfo {
	return groupInfo{
		groupName: groupName,
		members:   members,
	}
}

// 팀에 속한 멤버의 이름을 사전순으로 출력하는 메서드
func (g groupInfo) printMembers(writer *bufio.Writer) {
	sort.Strings(g.members)
	for _, member := range g.members {
		fmt.Fprintln(writer, member)
	}
}

func runQuiz(reader *bufio.Reader, writer *bufio.Writer, n int, m int) {
	groupInfos := make(map[string]groupInfo)
	memberInfos := make(map[string]string)

	// 데이터 셋팅
	for i := 0; i < n; i++ {
		var groupName string
		var memberCount int
		fmt.Fscanln(reader, &groupName)
		fmt.Fscanln(reader, &memberCount)

		members := make([]string, memberCount)
		for j := 0; j < memberCount; j++ {
			fmt.Fscanln(reader, &members[j])
			memberInfos[members[j]] = groupName
		}
		groupInfos[groupName] = newGroupInfo(groupName, members)
	}

	// 출력 조건
	for i := 0; i < m; i++ {
		var quiz string
		var quizType int
		fmt.Fscanln(reader, &quiz)
		fmt.Fscanln(reader, &quizType)

		if quizType == 0 {
			groupInfos[quiz].printMembers(writer)
		} else if quizType == 1 {
			fmt.Fprintln(writer, memberInfos[quiz])
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	runQuiz(reader, writer, n, m)
}
