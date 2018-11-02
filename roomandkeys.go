package main

type stack struct {
	data []int
}

func NewStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *stack) Pop() int {
	var ret int
	if s.Len() < 1 {
		ret = -1
	} else {
		ret = s.data[s.Len()-1]
		s.data = s.data[:s.Len()-1]
	}
	return ret
}

func (s *stack) Len() int {
	return len(s.data)
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	for k := range visited {
		visited[k] = false
	}
	visited[0] = true
	s := NewStack()
	s.Push(0)
	for s.Len() > 0 {
		for _, v := range rooms[s.Pop()] {
			if !visited[v] {
				s.Push(v)
				visited[v] = true
			}
		}
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true 
}
