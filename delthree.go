package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var ints []int
	for i := 0; i < n; i++ {
		tmp := 1
		ints = append(ints, tmp)
	}
	remian := n
	cnt := 0
	for i := 0; ; i++ {
		if remian == 1 {
			break
		}
		if ints[i%n] == 1 {
			cnt++
			if cnt == 3 {
				ints[i%n] = 0
				cnt = 0
				remian--
			}
		}
	}
	for k, v := range ints {
		if v == 1 {
			fmt.Printf("%d", k+1)
			break
		}
	}
}
