// TODO: 나이순 정렬
// https://www.acmicpc.net/problem/10814

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Member struct {
	age  int
	name string
	id   int // 가입 순서를 나타내기 위한 필드
}

type ByAge []Member

func (a ByAge) Len() int { return len(a) }

func (a ByAge) Less(i, j int) bool { // 나이가 같으면 가입 순으로, 아니면 나이 순으로 정렬
	if a[i].age == a[j].age {
		return a[i].id < a[j].id
	}
	return a[i].age < a[j].age
}

func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	members := make([]Member, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		temp := bufio.NewScanner(strings.NewReader(input))
		temp.Split(bufio.ScanWords)

		temp.Scan()
		age, _ := strconv.Atoi(temp.Text())

		temp.Scan()
		name := temp.Text()

		members[i] = Member{age, name, i}
	}

	sort.Sort(ByAge(members))

	for _, member := range members {
		fmt.Printf("%d %s\n", member.age, member.name)
	}
}
