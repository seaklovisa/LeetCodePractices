package main

import "sort"

func main() {
	println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
	println(countDays(10, [][]int{{2, 4}, {1, 3}}))
	println(countDays(10, [][]int{{1, 6}}))
}

func countDays(days int, meetings [][]int) int {
	//先排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0]-meetings[j][0] < 0
	})

	//最新合併後的區間
	//freedays = days - (所有會議區間佔用的天數總和)
	//days - 最新合併後的區間 <= 0 則滿足天數 回應所缺天數
	//days - 最新合併後的區間 > 0 則未滿足天數 繼續合併下一個

	curStart := meetings[0][0]
	curEnd := meetings[0][1]
	usedDays := 0

	for _, next := range meetings[1:] {
		nextStart := next[0]
		nextEnd := next[1]

		if nextEnd <= curEnd {
			curEnd = max(curEnd, nextEnd)
		} else {
			usedDays += curEnd - curStart + 1
			curStart = nextStart
			curEnd = nextEnd
		}
	}

	usedDays += curEnd - curStart + 1

	return days - usedDays

}
