/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
Note that you cannot sell a stock before you buy one
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bf := bufio.NewReader(os.Stdin)
	s, _, _ := bf.ReadLine()
	strs := strings.Split(string(s), ",")
	var ints []int
	for _, v := range strs {
		i, _ := strconv.Atoi(v)
		ints = append(ints, i)
	}
	var max = 0
	find2(ints, &max)
	fmt.Print(max)
}

// 暴力减法 n^n
func find1(ints []int, max *int) {
	for i := 0; i < len(ints); i++ {
		var temp = 0
		for j := i; j < len(ints); j++ {
			diff := ints[j] - ints[i]
			if temp < diff {
				temp = diff
			}
		}
		if *max < temp {
			*max = temp
		}
	}
}

// 历史记录法
func find2(ints []int, max *int) {
	tmpMin, tmpMax := ints[0], 0
	for i := 1; i < len(ints); i++ {
		if tmpMin > ints[i] {
			tmpMin = ints[i]
		}
		if ints[i]-tmpMin > tmpMax {
			tmpMax = ints[i] - tmpMin
		}
	}
	*max = tmpMax
}
