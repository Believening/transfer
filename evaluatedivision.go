/*
https://leetcode.com/problems/evaluate-division/description/

Equations are given in the format A / B = k, where A and B are variables represented as strings,
and k is a real number (floating point number). Given some queries, return the answers.
If the answer does not exist, return -1.0.

Given a / b = 2.0, b / c = 3.0.
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .
return [6.0, 0.5, -1.0, 1.0, -1.0 ].

 */
package main

import "fmt"

type strStack struct {
	data []string
}

func newStrStack() *strStack {
	return &strStack{
		data: make([]string, 0),
	}
}

func (s *strStack) push(str string) {
	s.data = append(s.data, str)
}

func (s *strStack) pop() string {
	var ret string
	if s.len() < 1 {
		ret = ""
	} else {
		ret = s.data[s.len()-1]
		s.data = s.data[:s.len()-1]
	}
	return ret
}

func (s *strStack) len() int {
	return len(s.data)
}

func findPath(g map[string]map[string]float64, src, dst string, visited []string) []string {
	if _,exist := g[src];!exist{
		return nil
	}
	// 初始化 本次路径
	path := make([]string, 0)
	path = append(path, src)
	if src == dst {
		return path
	}
	// 更新 前驱路径
	if visited == nil {
		visited = make([]string, 0)
	}
	visited = append(visited, src)
	for k := range g[src] {
		// 判断 该节点是否可用
		isVisited := false
		for _, v := range visited {
			if k == v {
				isVisited = true
			}
		}
		if isVisited {
			continue
		}
		// 已达 目的
		if k == dst {
			path = append(path, k)
			return path
		}
		// 备份 前驱路径 递归查找
		tmpV := make([]string, 0)
		tmpV = append(tmpV, visited...)
		tmp := findPath(g, k, dst,tmpV)
		if tmp != nil {
			path = append(path, tmp...)
			return path
		}
	}
	return nil
}

func calcOnePath(g map[string]map[string]float64, path []string) float64{
	if path == nil {
		return -1.0
	}
	if len(path) == 1{
		return 1.0
	}
	sum := 1.0
	for i:=0;i<len(path)-1;i++{
		sum *= g[path[i]][path[i+1]]
	}
	return sum
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := make(map[string]map[string]float64)
	for k, edge := range equations {
		if _, exit := g[edge[0]]; !exit {
			g[edge[0]] = make(map[string]float64, 0)
			g[edge[0]][edge[0]] = 0
		}
		if _, exit := g[edge[1]]; !exit {
			g[edge[1]] = make(map[string]float64, 0)
			g[edge[1]][edge[1]] = 0
		}
		g[edge[0]][edge[1]] = values[k]
		g[edge[1]][edge[0]] = 1 / values[k]
	}
	ret := make([]float64,0)
	for _,query := range queries{
		path:=findPath(g,query[0],query[1],nil)
		ret = append(ret,calcOnePath(g,path))
	}
	return ret
}

func main() {
	var equations [][]string = [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	g := make(map[string]map[string]float64)
	for _, edge := range equations {
		if _, exit := g[edge[0]]; !exit {
			g[edge[0]] = make(map[string]float64, 0)
			g[edge[0]][edge[0]] = 1.0
		}
		if _, exit := g[edge[1]]; !exit {
			g[edge[1]] = make(map[string]float64, 0)
			g[edge[1]][edge[1]] = 1.0
		}
		g[edge[0]][edge[1]] = 1.0
		g[edge[1]][edge[0]] = 1.0
	}
	p := findPath(g, "a", "c",nil)
	fmt.Println(p)
	p = findPath(g, "c", "a",nil)
	fmt.Println(p)

}

func dfs(g [][]int, s int, visited []int) {
	visited[s] = 1
	for _, v := range g[s] {
		if visited[v] != 1 {
			dfs(g, v, visited)
		}
	}
}
