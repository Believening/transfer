/*

Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
 */

package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var ret *TreeNode
	lenth := len(nums)
	switch lenth {
	case 0:
		ret = nil
	case 1:
		ret = &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	default:
		maxIdx := findMaxIdx(nums)
		ret = &TreeNode{
			Val:   nums[maxIdx],
			Left:  constructMaximumBinaryTree(nums[0:maxIdx]),
			Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
		}
	}
	return ret
}

func findMaxIdx(nums []int) int {
	var idx, max int
	idx, max = 0, nums[0]
	for i, v := range nums {
		if v > max {
			max = v
			idx = i
		}
	}
	return idx
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	_ = constructMaximumBinaryTree(nums)
}
