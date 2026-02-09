package main

import "sort"

func main() {
	println(countDays3(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
	println(countDays3(5, [][]int{{2, 4}, {1, 3}}))
	println(countDays3(6, [][]int{{1, 6}}))
}

func countDays3(days int, meetings [][]int) int {
	//1,3 5,7 9,10
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	//1
	curStart := meetings[0][0]
	//3
	curEnd := meetings[0][1]
	usedDays := 0

	for _, next := range meetings[1:] {
		//5
		nextStart := next[0]
		//7
		nextEnd := next[1]
		//有重疊，做合併區間
		if curEnd >= nextStart {
			curEnd = max(curEnd, nextEnd)
		} else {
			//沒有重疊 結算區間
			usedDays += curEnd - curStart + 1
			//5
			curStart = nextStart
			//7
			curEnd = nextEnd
		}
	}

	usedDays += curEnd - curStart + 1

	return days - usedDays

}
