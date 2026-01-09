package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{{4, 7}, {1, 4}}
	fmt.Println(merge(input))
}

/*
[1,3],[2,6],[8,10],[15,18]

1. rest={1,3},last={1,3},cur={2,6},
cur[0]=2 <= last[1]=3 有重疊,cur[1]=6 > last[1]=3 last[1]更新為6 => last={1,6}
2.cur={8,10},last={1,6},cur[0]=8 <= last[1]=6 沒有重疊,rest={1,6},{8,10}
3.cur={15,18},last從rest取最後一個 = {8,10},
cur[0]=15 <= last[1]=10 沒有重疊,rest={1,6},{8,10},{15,18}
4.回傳結果 {1,6},{8,10},{15,18}ß
*/

func merge(intervals [][]int) [][]int {
	//先依左邊元素排序，決定左邊界
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//比較元素

	//合併後的結果存放在這，先取第一個
	rest := [][]int{intervals[0]}

	//從下一個元素開始比較
	for _, cur := range intervals[1:] {
		//取最後一個合併結果，與下一個元素比較
		last := rest[len(rest)-1]

		//有沒有重疊
		if cur[0] <= last[1] {
			//決定右邊界哪個大
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			//沒有重疊 加入結果
			rest = append(rest, cur)
		}
	}
	return rest
}
