package main

import "fmt"

func quicksort(nums []int, left, right int) {
	if left >= right {
		return
	}
	if left+1 == right {
		nums[left], nums[right] = nums[right], nums[left]
		return
	}
	flag := (left + right) / 2
	nums[left], nums[flag] = nums[flag], nums[left]
	i, j := left+1, right
	for {
		for nums[i] < nums[left] {
			i++
		}
		for nums[j] > nums[left] {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			break
		}
	}
	if i == j {
		if nums[i] < nums[left] {
			nums[left], nums[i] = nums[i], nums[left]
		}
		j--
	}
	quicksort(nums, left, j)
	quicksort(nums, i, right)
}

func main() {
	var ints = []int{3, 4, 7, 2, 1, 8, 9}
	fmt.Println(ints)
	quicksort(ints, 0, len(ints)-1)
	fmt.Println(ints)
}
