package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func langestCommomSequenceWithTwo(a, b []byte) int {
	var dp [][]int
	dp = make([][]int, len(a))
	for k := range dp {
		dp[k] = make([]int, len(b))
	}
	for k := range a{
		if b[0] == a[k]{
			dp[k][0] = 1
		}
	}
	for k := range b{
		if a[0] == b[k]{
			dp[0][k] = 1
		}
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if a[i] == b[j] {
				dp[i][j] = max(dp[i][j-1]+1, dp[i-1][j]+1)
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(a)-1][len(b)-1]
}

func main() {
	var n int
	var strs []string
	fmt.Scanf("%d", &n)
	for n > 0 {
		var s string
		fmt.Scanf("%s", &s)
		strs = append(strs, s)
		n--
	}
	l := langestCommomSequenceWithTwo([]byte(strs[0]), []byte(strs[1]))
	fmt.Println(l)

}
