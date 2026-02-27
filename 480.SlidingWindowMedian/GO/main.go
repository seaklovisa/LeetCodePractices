package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(medianSlidingWindow(nums, k))
	// expected: [1 -1 -1 3 5 6]

}

/*------------ Heaps -------------*/
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }
func (h MinHeap) Top() int           { return h[0] }

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }
func (h MaxHeap) Top() int           { return h[0] }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type DualHeap struct {
	small *MaxHeap
	large *MinHeap
	//數應該被刪除幾次
	del map[int]int //lazy delete
	k   int
	//smallSize / largeSize 是有效元素數(不含被del標記)
	smallSize int
	largeSize int
}

func newDualHeap(k int) *DualHeap {
	s := &MaxHeap{}
	l := &MinHeap{}
	heap.Init(s)
	heap.Init(l)
	return &DualHeap{
		small:     s,
		large:     l,
		del:       make(map[int]int),
		k:         k,
		smallSize: 0,
		largeSize: 0,
	}
}

func (dh *DualHeap) Insert(x int) {
	if dh.small.Len() == 0 || x <= dh.small.Top() {
		heap.Push(dh.small, x)
		dh.smallSize++
	} else {
		heap.Push(dh.large, x)
		dh.largeSize++
	}

	dh.rebalance()
}

// 邏輯刪除
func (dh *DualHeap) Erase(x int) {
	dh.del[x]++

	if x <= dh.small.Top() {
		dh.smallSize--
		if dh.small.Len() > 0 && x == dh.small.Top() {
			dh.pruneSmall()
		}
	} else {
		dh.largeSize--
		if dh.large.Len() > 0 && x == dh.large.Top() {
			dh.pruneLarge()
		}
	}
	dh.rebalance()
}

func (dh *DualHeap) rebalance() {
	// 先把可能會用到 top 的堆清乾淨（只為了後面 pop 正確）
	dh.pruneSmall()
	dh.pruneLarge()

	if dh.smallSize > dh.largeSize+1 {
		// small.top 要搬到 large
		x := heap.Pop(dh.small).(int)
		dh.smallSize--
		heap.Push(dh.large, x)
		dh.largeSize++
	} else if dh.smallSize < dh.largeSize {
		// large.top 要搬到 small
		x := heap.Pop(dh.large).(int)
		dh.largeSize--
		heap.Push(dh.small, x)
		dh.smallSize++
	}

	// 收尾：確保 rebalance 後 top 乾淨，median 取值安全
	dh.pruneSmall()
	dh.pruneLarge()
}

func (dh *DualHeap) pruneSmall() {
	for dh.small.Len() > 0 {
		x := dh.small.Top()
		if c, ok := dh.del[x]; ok && c > 0 {
			heap.Pop(dh.small)
			if c == 1 {
				delete(dh.del, x)
			} else {
				dh.del[x] = c - 1
			}
		} else {
			break
		}
	}
}

func (dh *DualHeap) pruneLarge() {
	for dh.large.Len() > 0 {
		x := dh.large.Top()
		if c, ok := dh.del[x]; ok && c > 0 {
			heap.Pop(dh.large)
			if c == 1 {
				delete(dh.del, x)
			} else {
				dh.del[x] = c - 1
			}
		} else {
			break
		}
	}
}

// GetMedian: 取 median 前要確保 top 都是乾淨的（pruned）
// 但我們在 insert/erase/rebalance 內都會 prune，所以這裡通常已經乾淨
func (dh *DualHeap) GetMedian() float64 {
	// 安全起見：再次 prune 一下 top（不貴）
	dh.pruneSmall()
	dh.pruneLarge()

	if dh.k%2 == 1 {
		return float64(dh.small.Top())
	}
	return (float64(dh.small.Top()) + float64(dh.large.Top())) / 2.0
}

func medianSlidingWindow(nums []int, k int) []float64 {
	dh := newDualHeap(k)
	res := make([]float64, 0, max(0, len(nums)-k+1))

	for i := 0; i < len(nums); i++ {
		dh.Insert(nums[i])

		if i >= k {
			dh.Erase(nums[i-k])
		}

		if i >= k-1 {
			res = append(res, dh.GetMedian())
		}
	}
	return res
}
