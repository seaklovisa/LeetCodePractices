package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mostBooked(1, [][]int{
		{0, 5},
		{6, 10},
		// {3, 5},
		// {4, 9},
		// {6, 8},
	}))
}

// ----------  空房間的MinHeap 拿最小的空房間id----------
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

// ---------- busy rooms 存(endtime,roomid) 拿最早結束的房間 ----------
type BusyRoom struct {
	endTime int
	roomId  int
}

type BusyMinHeap []BusyRoom

func (h BusyMinHeap) Len() int { return len(h) }
func (h BusyMinHeap) Less(i, j int) bool {
	if h[i].endTime != h[j].endTime {
		return h[i].endTime < h[j].endTime
	}

	return h[i].roomId < h[j].roomId
}
func (h BusyMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *BusyMinHeap) Push(x any)   { *h = append(*h, x.(BusyRoom)) }
func (h *BusyMinHeap) Pop() any     { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

// LeetCode 2402. Meeting Rooms III
// 思路：兩個 min-heap
// 1) available：存「空房間的 roomId」，永遠拿最小 roomId
// 2) busy：存「(endTime, roomId)」，永遠拿最早結束的房間
//
// 規則：
//   - 會議依 start 由小到大處理
//   - 每來一場會議：先把 busy 中 endTime <= start 的房間釋放回 available
//   - 若有空房：用最小 roomId，會議照原本時間跑
//   - 若沒空房：取最早結束的房間 (endTime 最小；若同 endTime 取 roomId 最小)
//     會議延後到 endTime 開始，長度不變：newEnd = endTime + (end-start)
//   - 統計每個 room 使用次數，最後回傳使用最多次的最小 roomId
func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	availableRoom := &MinHeap{}
	for i := 0; i < n; i++ {
		*availableRoom = append(*availableRoom, i)
	}
	heap.Init(availableRoom)

	busy := &BusyMinHeap{}
	heap.Init(busy)

	count := make([]int, n)

	for _, mt := range meetings {
		start := mt[0]
		end := mt[1]
		duration := end - start

		//釋出房間
		for busy.Len() > 0 && (*busy)[0].endTime <= start {
			top := heap.Pop(busy).(BusyRoom)
			heap.Push(availableRoom, top.roomId)
		}

		var room int
		var actualEnd int

		if availableRoom.Len() > 0 {
			room = heap.Pop(availableRoom).(int)
			actualEnd = end
		} else {
			top := heap.Pop(busy).(BusyRoom)
			room = top.roomId
			actualEnd = top.endTime + duration
		}

		count[room]++
		heap.Push(busy, BusyRoom{endTime: actualEnd, roomId: room})
	}

	bestRoom := 0
	for i := 1; i < n; i++ {
		if count[i] > count[bestRoom] {
			bestRoom = i
		}
	}

	return bestRoom
}
