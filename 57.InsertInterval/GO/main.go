package main

import (
	"fmt"
)

func main() {
	newInterval := []int{4, 8}
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	ns, ne := newInterval[0], newInterval[1]

	for _, it := range intervals {
		s, e := it[0], it[1]

		if e < ns {
			res = append(res, it)
		} else if s > ne {
			res = append(res, []int{ns, ne})
			ns, ne = s, e
		} else {
			if s < ns {
				ns = s
			}
			if e > ne {
				ne = e
			}
		}
	}

	res = append(res, []int{ns, ne})
	return res
}
