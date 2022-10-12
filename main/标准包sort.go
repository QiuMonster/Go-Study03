package main

import (
	"fmt"
	"sort"
)

func go_sort() {
	s := []int{1, 2, 3, 6, 3, 8}
	fmt.Println("排序前的数组:", s)
	sort.Ints(s)
	fmt.Println("排序后的数组:", s)
}
