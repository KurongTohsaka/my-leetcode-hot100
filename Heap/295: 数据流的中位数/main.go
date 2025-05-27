package main

import (
	"container/heap"
	"sort"
)

func main() {

}

type MedianFinder struct {
	left  hp // 入堆的元素取相反数，变成最大堆
	right hp // 最小堆
}

func Constructor() (_ MedianFinder) { return }

func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == mf.right.Len() {
		heap.Push(&mf.right, num)
		heap.Push(&mf.left, -heap.Pop(&mf.right).(int))
	} else {
		heap.Push(&mf.left, -num)
		heap.Push(&mf.right, -heap.Pop(&mf.left).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() > mf.right.Len() {
		return float64(-mf.left.IntSlice[0])
	}
	return float64(mf.right.IntSlice[0]-mf.left.IntSlice[0]) / 2.0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
