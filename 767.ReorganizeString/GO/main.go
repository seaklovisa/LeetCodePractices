package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(reorganizeString("aab"))
}

func reorganizeString(s string) string {
	//先統計各個字元的出現次數，轉換成map[ch]=count
	counter := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	//找出最多個數字元的數
	maxCount := 0
	for _, cnt := range counter {
		if cnt > maxCount {
			maxCount = cnt
		}
	}

	//k = 最多字元的計數，n = 字串，n-k剩下的字串，k-1 至少需要這些數量的字元插入，才能不連續
	//n-k >= k-1 => 2k <= n +1 => k <= 1/2(n+1)
	//無法形成不連續的字串
	if maxCount > (len(s)+1)/2 {
		return ""
	}

	//將map的每一項推進maxheap
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for ch, cnt := range counter {
		heap.Push(maxHeap, Item{ch: ch, count: cnt})
	}

	result := make([]byte, 0, len(s))

	//每次取pop兩個，append成string，次數減1，有剩就推回去maxheap
	for maxHeap.Len() >= 2 {
		first := heap.Pop(maxHeap).(Item)
		second := heap.Pop(maxHeap).(Item)

		first.count--
		second.count--

		result = append(result, first.ch, second.ch)

		if first.count > 0 {
			heap.Push(maxHeap, first)
		}
		if second.count > 0 {
			heap.Push(maxHeap, second)
		}
	}

	//剩下1個就append
	if maxHeap.Len() == 1 {
		last := heap.Pop(maxHeap).(Item)
		result = append(result, byte(last.ch))
	}

	return string(result)

}

/*宣告maxheap*/
type Item struct {
	ch    byte
	count int
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
