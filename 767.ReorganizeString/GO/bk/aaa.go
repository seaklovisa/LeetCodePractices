package aaaa

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(reorganizeString("aab"))
}

func reorganizeString(s string) string {
	if len(s) <= 1 {
		return s
	}

	freq := make(map[byte]int)
	maxFreq := 0
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
		if freq[s[i]] > maxFreq {
			maxFreq = freq[s[i]]
		}
	}

	if maxFreq > (len(s)+1)/2 {
		return ""
	}

	h := &maxHeap{}
	heap.Init(h)
	for ch, count := range freq {
		heap.Push(h, item{char: ch, count: count})
	}

	result := make([]byte, 0, len(s))
	for h.Len() >= 2 {
		first := heap.Pop(h).(item)
		second := heap.Pop(h).(item)

		result = append(result, first.char, second.char)

		first.count--
		second.count--

		if first.count > 0 {
			heap.Push(h, first)
		}
		if second.count > 0 {
			heap.Push(h, second)
		}
	}

	if h.Len() == 1 {
		last := heap.Pop(h).(item)
		result = append(result, last.char)
	}

	return string(result)
}

type item struct {
	char  byte
	count int
}

type maxHeap []item

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}
