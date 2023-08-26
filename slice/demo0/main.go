package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

var arr = [...]string{"aa", "bb", "cc", "dd", "DD", "dddd", "ee", "EE", "ff", "gg", "hh", "HH", "HH", "GG", "GG"}
var list = arr[2:6]

func main() {
	fmt.Println(slices.BinarySearch(list, "bb"))
	fmt.Println(slices.BinarySearchFunc(list, "bb", func(a, b string) int {
		return cmp.Compare(a, b)
	}))

	fmt.Println(list, len(list), cap(list))
	list1 := slices.Clip(list)
	fmt.Println(list1, len(list1), cap(list1))

	fmt.Println(slices.Clone(list))

	fmt.Println(slices.Compact(slices.Clone(arr[:])))
	fmt.Println(slices.CompactFunc(slices.Clone(arr[:]), func(s string, s2 string) bool {
		return s[:2] == s2[:2]
	}))
	fmt.Println(slices.CompactFunc(slices.Clone(arr[:]), func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	}))
	fmt.Println(slices.Contains(list, "dd"))
	fmt.Println(slices.ContainsFunc(list, func(s string) bool {
		return strings.ToLower(s) == "dd"
	}))
	fmt.Println(slices.Insert(list, 2, "ii", "jj"))
	fmt.Println(slices.Delete(list, 2, 3))
	fmt.Println(slices.DeleteFunc(list, func(s string) bool {
		return strings.ToUpper(s) == s
	}))
	fmt.Println(slices.Replace(arr[:], 1, 3, "aa", "bb"))
	fmt.Println(cap(slices.Grow([]int{1, 2, 3}, 15)))
	fmt.Println(slices.Equal(arr[:], arr[:]))
	fmt.Println(slices.Index(arr[:], "aa"))
	slices.Reverse(arr[:])
	fmt.Println(arr)
	fmt.Println(slices.IsSorted(arr[:]))
}
