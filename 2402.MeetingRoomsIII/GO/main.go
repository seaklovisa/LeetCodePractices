package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mostBooked())
}

// ---------- available rooms (min-heap by room index) ----------
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

// ---------- busy rooms (min-heap by endTime, tie by room index) ----------
func (h BusyHeap) Len() int { return len(h) }
func (h BusyHeap) Less(i, j int) bool {
	if h[i].end != h[j].end {
		return h[i].end < h[j].end
	}
	return h[i].room < h[j].room
}
func (h BusyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *BusyHeap) Push(x any)   { *h = append(*h, x.(Busy)) }
func (h *BusyHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type Busy struct {
	end  int
	room int
}
type BusyHeap []Busy

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] != meetings[j][0] {
			return meetings[i][0] < meetings[j][0]
		}
		return meetings[i][1] < meetings[j][1]
	})

	count := make([]int, n)

	avail := &IntHeap{}
	heap.Init(avail)
	for i := 0; i < n; i++ {
		heap.Push(avail, i)
	}

	busy := &BusyHeap{}
	heap.Init(busy)

	for _, mt := range meetings {
		start, end := mt[0], mt[1]
		dur := end - start

		// release rooms that have finished by 'start'
		for busy.Len() > 0 && (*busy)[0].end <= start {
			b := heap.Pop(busy).(Busy)
			heap.Push(avail, b.room)
		}

		var room int
		var realEnd int

		if avail.Len() > 0 {
			room = heap.Pop(avail).(int)
			realEnd = end
		} else {
			// delay to earliest finishing room
			b := heap.Pop(busy).(Busy)
			room = b.room
			start = b.end
			realEnd = start + dur
		}

		heap.Push(busy, Busy{end: realEnd, room: room})
		count[room]++
	}

	bestRoom := 0
	for i := 1; i < n; i++ {
		if count[i] > count[bestRoom] {
			bestRoom = i
		}
	}
	return bestRoom
}
