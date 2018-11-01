package main

import "fmt"

/*
建一个最大堆，然后将堆顶和最后一个元素换位置
重新建堆
原地排序
 */
func heapSort(ints []int) {
	s := len(ints) / 2 //最后一个有儿子
	for i := s; i > -1; i-- { //建堆
		heap(ints, i, len(ints)-1)
	}

	for i := len(ints) - 1; i > 0; i-- {
		ints[i], ints[0] = ints[0], ints[i]
		heap(ints, 0, i-1)
	}
}
func heap(ints []int, start, end int) { //up
	l := start*2 + 1
	if l > end {
		return
	}
	r := start*2 + 2
	p := l
	if r <= end && ints[l] > ints[r] {
		p = r
	}
	if ints[p] > ints[start] {
		return
	}
	ints[p], ints[start] = ints[start], ints[p]
	heap(ints, p, end)
}

func insert(ints []int, x int) []int {
	ints = append(ints, x)
	cur := len(ints) - 1
	for i := cur / 2; ints[i] > ints[cur]; i = i / 2 {
		ints[i], ints[cur] = ints[cur], ints[i]
		cur = i
	}
	return ints
}

func makeheap(ints []int) {
	s := len(ints) / 2
	for i := s; i > -1; i-- {
		heap(ints, i, len(ints)-1)
	}
}

func main() {
	nums := []int{3, 4, 78, 23, 4, 65, 23, 5, 45, 8, 2}
	t := []int{}
	for _, v := range nums {
		t = insert(t, v)
	}
	fmt.Println(t)
	makeheap(nums)
	fmt.Println(nums)
}
