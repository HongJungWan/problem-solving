// TODO: 문자열 내 마음대로 정렬하기
// https://school.programmers.co.kr/learn/courses/30/lessons/12915

package main

import (
	"fmt"
	"sort"
)

type ByIndex struct {
	strings []string
	index   int
}

func (byIndex ByIndex) Len() int {
	return len(byIndex.strings)
}

func (byIndex ByIndex) Swap(i, j int) {
	byIndex.strings[i], byIndex.strings[j] = byIndex.strings[j], byIndex.strings[i]
}

func (byIndex ByIndex) Less(i, j int) bool {
	// 주어진 인덱스의 문자가 같을 경우, 사전순으로 문자열 비교
	if byIndex.strings[i][byIndex.index] == byIndex.strings[j][byIndex.index] {
		return byIndex.strings[i] < byIndex.strings[j]
	}

	// 주어진 인덱스의 문자가 다를 경우, 해당 문자를 비교
	return byIndex.strings[i][byIndex.index] < byIndex.strings[j][byIndex.index]
}

func sortStringsByIndex(strings []string, index int) []string {
	sort.Sort(ByIndex{strings, index})
	return strings
}

func main() {
	strings1 := []string{"sun", "bed", "car"}
	sortedStrings1 := sortStringsByIndex(strings1, 1)
	fmt.Println(sortedStrings1) // Output: ["car", "bed", "sun"]

	strings2 := []string{"abce", "abcd", "cdx"}
	sortedStrings2 := sortStringsByIndex(strings2, 2)
	fmt.Println(sortedStrings2) // Output: ["abcd", "abce", "cdx"]
}

/*
사용자 정의 정렬을 구현할 때, sort.Interface를 구현해야 한다.
이를 위해 Len(), Less(), Swap() 메서드를 정의해야 한다.

byIndex.strings[i][byIndex.index]는 i번째 문자열의 byIndex.index 인덱스에 있는 문자를 나타낸다.
byIndex.strings[j][byIndex.index]는 j번째 문자열의 byIndex.index 인덱스에 있는 문자를 나타낸다.
*/
