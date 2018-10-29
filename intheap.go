package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *intHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func main() {
	hd := []int{9, 2, 3, 5, 78, 2, 3, 4, 6, 1}
	h := new(intHeap)
	for _, v := range hd {
		h.Push(v)
	}
	fmt.Println(h)
	heap.Init(h)
	fmt.Println(h)
	i := heap.Pop(h)
	fmt.Println(i, h)
	j := heap.Remove(h, 7)
	fmt.Println(j, h)

}
