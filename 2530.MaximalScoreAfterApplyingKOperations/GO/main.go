package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	//nums塞進maxheap
	for _, ele := range nums {
		heap.Push(maxHeap, ele)
	}

	result := int64(0)
	//for k 次
	//x = 彃出
	//p = 計算ceil(x/3)
	//result += x
	//p推入maxheap
	for k > 0 {
		x := heap.Pop(maxHeap).(int)
		p := (x + 2) / 3
		result += int64(x)
		heap.Push(maxHeap, p)
		k--
	}

	return result
}
