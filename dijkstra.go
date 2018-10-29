package main

import (
	"fmt"
	"math"
)

func dijkstra(g [][]int, s int) {
	var d, Q, S,P []int
	d = make([]int, len(g[s]))
	for k, v := range g[s] {
		d[k] = v
	}
	Q = make([]int, len(g[s]))
	for k := range Q {
		Q[k] = k
	}

	defer func() {
		fmt.Println(P)
		fmt.Println(d)
	}()

	P = make([]int, len(g[s]))
	for len(Q) > 0 {
		// find min Q
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
			if g[u][v] > 0 {
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
		{0, 2, 8, 1, -1, -1, -1, -1},
		{2, 0, 6, 7, 1, 2, -1, -1},
		{8, 6, 0, -1, 4, -1, 2, -1},
		{1, 7, -1, 0, -1, -1, 9, -1},
		{-1, 1, 4, -1, 0, 3, -1, 9},
		{-1, 2, -1, 0, 3, 0, 4, 6},
		{-1, -1, 2, 9, -1, 4, 0, 2},
		{-1, -1, -1, -1, 9, 6, 2, 0},
	}
	dijkstra(ints, 0)
}
