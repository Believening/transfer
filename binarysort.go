package main

import "fmt"

type node struct {
	val   int
	left  *node
	rigrt *node
}

func newNode() *node {
	return &node{
		val: -1,
	}
}

func Insert(root *node, x int) {
	n := root
	if n.val == -1 {
		n.val = x
		return
	}
	if x > n.val {
		if n.rigrt == nil{
			n.rigrt = newNode()
		}
		n = n.rigrt
		Insert(n, x)
	}
	if x < n.val {
		if n.left == nil{
			n.left = newNode()
		}
		n = n.left
		Insert(n, x)
	}
	return
}

func main() {
	ints := []int{2, 1, 3}
	t := newNode()
	for _, v := range ints {
		Insert(t, v)
	}
	fmt.Println(t)

}
