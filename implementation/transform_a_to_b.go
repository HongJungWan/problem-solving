package main

import (
	"fmt"
	"sort"
	"strings"
)

func transformAtoB(before string, after string) int {
	sortedBefore := sortString(before)
	sortedAfter := sortString(after)

	if sortedBefore == sortedAfter {
		return 1
	} else {
		return 0
	}
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func main() {
	before1 := "olleh"
	after1 := "hello"

	before2 := "allpe"
	after2 := "apple"

	fmt.Println(transformAtoB(before1, after1)) // 1
	fmt.Println(transformAtoB(before2, after2)) // 0
}
