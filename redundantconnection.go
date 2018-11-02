/*
https://leetcode.com/problems/redundant-connection/description/
 */

package main

import (
	"fmt"
)

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	set := make([]int, n+1)
	for k := range set {
		set[k] = -1
	}
	findRoot := func(set []int, idx int) int {
		for {
			if set[idx] == -1 || set[idx] == idx {
				break
			}
			idx = set[idx]
		}
		return idx
	}

	unionRoot := func(set []int, x, y int) {
		rootX := findRoot(set, x)
		rootY := findRoot(set, y)
		set[rootY] = rootX
	}

	ret := make([]int, 0)
	for _, edge := range edges {
		switch {
		case set[edge[0]] == -1 && set[edge[1]] == -1:
			set[edge[0]] = edge[0]
			set[edge[1]] = edge[0]
		case set[edge[0]] == -1 && set[edge[1]] != -1:
			set[edge[0]] = edge[1]
		case set[edge[0]] != -1 && set[edge[1]] == -1:
			set[edge[1]] = edge[0]
		case set[edge[0]] != -1 && set[edge[1]] != -1:
			if findRoot(set, edge[0]) == findRoot(set, edge[1]) {
				ret = append(ret, edge...)
			} else {
				unionRoot(set, edge[0], edge[1])
			}
		}
	}
	return ret
}

type simpleUnionSet struct {
	set []int
}

func (s *simpleUnionSet) find(x int) int {
	for s.set[x] != x {
		x = s.set[x]
		// s.set[x] = s.find(s.set[x]) // 将层数压缩
	}
	return x
}

func (s *simpleUnionSet) union(x, y int) {
	s.set[s.find(y)] = s.find(x)
}

func main() {
	ret := findRedundantConnection([][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	})
	fmt.Println(ret)
}
