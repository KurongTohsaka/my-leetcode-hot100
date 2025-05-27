package main

import "math"

func main() {

}

func findMin(nums []int) int {
	minVal := math.MinInt
	reverse := -1
	// 找到反转位置，reverse 为反转后的后半段序列的首部
	for i := 0; i < len(nums); i++ {
		if nums[i] > minVal {
			minVal = nums[i]
		} else {
			reverse = i
			break
		}
	}

	ans := math.MinInt
	if reverse == -1 {
		ans = nums[0]
	} else {
		ans = nums[reverse]
	}
	return ans
}

func findMin2(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] {
			// 在循环的每一步中，最小值要么是 right 指向的元素，要么在 right 的左边（如果 right 移动了）
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right]
}
