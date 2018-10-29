package main

import (
	"container/list"
	"fmt"
)

type keys struct {
	room int
}

func main() {
	l := list.New()
	// 队列操作
	l.PushBack(1) //入队
	l.PushBack(2)
	ef := l.Front()
	l.Remove(ef) //出队
	fmt.Println(ef.Value)
	fmt.Println(l)
	// 栈操作
	l.PushBack(3) //入栈
	eb := l.Back()
	l.Remove(eb) //出栈
	fmt.Println(eb.Value)
	fmt.Println(l)
}
