package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type Item struct {
	ele int
	ct  int
}

type MaxHeap []Item

func (h MaxHeap) Len() int      { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].ct > h[j].ct
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	counter := make(map[int]int)

	//計算每個元素的出現次數
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	//塞進maxheap
	for ele, count := range counter {
		heap.Push(maxHeap, Item{ele: ele, ct: count})
	}

	//依序彈出
	result := make([]int, 0, k)
	index := k
	for index > 0 {
		result = append(result, heap.Pop(maxHeap).(Item).ele)
		index--
	}

	return result
}
