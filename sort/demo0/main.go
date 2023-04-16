package main

import (
	"fmt"
	"sort"
)

type People struct {
	Name string
	Age  int
}

type sortType int

const (
	sortAsc sortType = iota
	sortDesc
)

type lessFunc func(s1, s2 *People) bool

type sortLess struct {
	Less     lessFunc
	SortType sortType
}

type multiSorter struct {
	datas []People
	less  []sortLess
}

func (ms *multiSorter) Sort(data []People) {
	ms.datas = data
	sort.Sort(ms)
}

func (ms *multiSorter) Len() int {
	return len(ms.datas)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.datas[i], ms.datas[j] = ms.datas[j], ms.datas[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.datas[i], &ms.datas[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less.Less(p, q):
			return less.SortType == sortAsc
		case less.Less(q, p):
			return less.SortType == sortDesc
		}
	}
	return ms.less[k].Less(p, q) && ms.less[k].SortType == sortAsc
}

func main() {
	var (
		peoples       []People
		byName, byAge func(p1, p2 *People) bool
	)
	peoples = []People{
		{Name: "alice", Age: 19},
		{Name: "alice", Age: 18},
		{Name: "alice", Age: 18},
		{Name: "bob", Age: 20},
		{Name: "bob", Age: 21},
		{Name: "bob", Age: 20},
		{Name: "cindy", Age: 18},
	}
	byName = func(p1, p2 *People) bool {
		return p1.Name < p2.Name
	}
	byAge = func(p1, p2 *People) bool {
		return p1.Age < p2.Age
	}
	sort.Sort(&multiSorter{
		datas: peoples,
		less:  []sortLess{{Less: byName, SortType: sortDesc}, {Less: byAge, SortType: sortAsc}},
	})

	fmt.Println(peoples)

}
