/*
Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number on it represented by array nums.
You are asked to burst all the balloons. If the you burst balloon i you will get nums[left] * nums[i] * nums[right] coins.
Here left and right are adjacent indices of i. After the burst, the left and right then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMaxCoins(dp [][]int, nums []int, left, right int) int {
	if left+1 == right {
		return 0
	}
	if dp[left][right] > 0 {
		return dp[left][right]
	}
	var max int = 0
	for i := left + 1; i < right; i++ {
		temp := nums[left]*nums[i]*nums[right] + getMaxCoins(dp, nums, left, i) + getMaxCoins(dp, nums, i, right)
		if max < temp {
			max = temp
		}
	}
	dp[left][right] = max
	return dp[left][right]
}

func main() {
	var ints []int
	r := bufio.NewReader(os.Stdin)
	sInts, _, _ := r.ReadLine()
	sBytes := strings.Fields(string(sInts))
	for _, v := range sBytes {
		if i, _ := strconv.Atoi(v); i > 0 {
			ints = append(ints, i)
		}
	}
	ints = append([]int{1}, append(ints, 1)...)
	var dp [][]int
	for i := 0; i < len(ints); i++ {
		temp := make([]int, len(ints))
		dp = append(dp, temp)
	}
	fmt.Println(getMaxCoins(dp, ints, 0, len(ints)-1))
}
