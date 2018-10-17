package main

import (
	"fmt"
)

var (
	dp  [][]int
	d   []int
	str []byte
)

func isPalindomic(i, j int) int {
	if dp[i][j] == -1 {
		switch j - i {
		case 0:
			dp[i][j] = 1
		case 1:
			if str[i] == str[j] {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		default:
			if isPalindomic(i+1, j-1) == 1 && str[i] == str[j] {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[i][j]
}

func isPalindomic2(i, j int) bool {
	if d[i] < j {
		switch j - i {
		case 0:
			d[i] = j
		case 1:
			if str[i] == str[j] {
				d[i] = j
			}
		default:
			if isPalindomic2(i+1, j-1) && str[i] == str[j] {
				d[i] = j
			}
		}
	}
	return d[i] == j
}

func main() {
	fmt.Scanf("%s", &str)
	//for i := 0; i < len(str); i++ {
	//	temp := make([]int, len(str))
	//	for idx, _ := range temp {
	//		temp[idx] = -1
	//	}
	//	dp = append(dp, temp)
	//}
	l := 1
	var start, end int
	d = make([]int, len(str))
	for idx, _ := range d {
		d[idx] = -1
	}
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if isPalindomic2(i, j) {
				if l <= j-i+1 {
					l = j - i + 1
					start = i
					end = j
				}
			}
		}
	}
	fmt.Println(l)
	fmt.Printf("%s", str[start:end+1])
}
