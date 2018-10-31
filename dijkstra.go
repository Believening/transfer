package main

import (
	"fmt"
	"math"
)

func dijkstra(g [][]int, s int) {
	var d, Q, S, P []int
	d = make([]int, len(g[s]))
	P = make([]int, len(g[s]))
	for k, v := range g[s] {
		d[k] = v
		P[k] = s
	}
	Q = make([]int, len(g[s]))
	for k := range Q {
		Q[k] = k
	}

	defer func() {
		fmt.Println(P)
		fmt.Println(d)
	}()

	for len(Q) > 0 {
		// find min Q
		// 此处使用堆实现 Q 每次取堆顶就可以了
		min, key, u := math.MaxInt64, -1, -1
		for k, v := range Q {
			if d[v] >= 0 && d[v] < min {
				min = d[v]
				key = k
				u = v
			}
		}
		switch {
		case key < 0:
			return
		case key+1 == len(Q):
			Q = Q[:key]
		default:
			Q = append(Q[:key], Q[key+1:]...)
		}

		// update S and D
		S = append(S, u)
		for _, v := range Q {
			if g[u][v] >= 0 {
				if d[v] < 0 {
					d[v] = g[u][v] + d[u]
					P[v] = u
				} else {
					if g[u][v]+d[u] < d[v] {
						d[v] = g[u][v] + d[u]
						P[v] = u
					}
				}
			}
		}
	}
}

func main() {
	var ints = [][]int{
		{3, 5, 78},
		{2, 1, 1},
		{1, 3, 0},
		{4, 3, 59},
		{5, 3, 85},
		{5, 2, 22},
		{2, 4, 23},
		{1, 4, 43},
		{4, 5, 75},
		{5, 1, 15},
		{1, 5, 91},
		{4, 1, 16},
		{3, 2, 98},
		{3, 4, 22},
		{5, 4, 31},
		{1, 2, 0},
		{2, 5, 4},
		{4, 2, 51},
		{3, 1, 36},
		{2, 3, 59},
	}
	//var ints = [][]int{
	//	{0,0,0,43,91},
	//	{1,0,59,23,4},
	//	{36,98,0,22,78},
	//	{16,51,59,0,75},
	//	{15,22,85,31,0},
	//}
	fmt.Println(networkDelayTime(ints, 5, 5))
}

func networkDelayTime(times [][]int, N int, K int) int {
	var g [][]int
	for i := 0; i < N; i++ {
		tmp := make([]int, N)
		for j := 0; j < N; j++ {
			if j == i {
				tmp[j] = 0
			} else {
				tmp[j] = -1
			}
		}
		g = append(g, tmp)
	}
	for _, v := range times {
		g[v[0]-1][v[1]-1] = v[2]
	}

	var d, Q, S []int
	d = make([]int, N)
	for k, v := range g[K-1] {
		d[k] = v
	}
	Q = make([]int, N)
	for k := range Q {
		Q[k] = k
	}

	defer func() {
		fmt.Println(d)
	}()

	for len(Q) > 0 {
		// find min Q
		// 此处使用堆实现 Q 每次取堆顶就可以了
		min, key, u := math.MaxInt64, -1, -1
		for k, v := range Q {
			if d[v] >= 0 && d[v] < min {
				min = d[v]
				key = k
				u = v
			}
		}
		switch {
		case key < 0:
			return -1
		case key+1 == len(Q):
			Q = Q[:key]
		default:
			Q = append(Q[:key], Q[key+1:]...)
		}

		// update S and D
		S = append(S, u)
		for _, v := range Q {
			if g[u][v] >= 0 {
				if d[v] < 0 {
					d[v] = g[u][v] + d[u]
				} else {
					if g[u][v]+d[u] < d[v] {
						d[v] = g[u][v] + d[u]
					}
				}
			}
		}
	}
	max := d[0]
	for _, v := range d {
		if v > max {
			max = v
		}
	}
	return max
}
