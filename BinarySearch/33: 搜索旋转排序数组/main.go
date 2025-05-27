package main

import "math"

func main() {

}

func search(nums []int, target int) int {
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

	bs := func(left, right int) int {
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	ans := -1
	// 寻找 target 在哪一部分中
	if reverse == -1 { // 没有反转，按照普通的二分进行
		ans = bs(0, len(nums)-1)
	} else if target >= nums[reverse] && target <= nums[len(nums)-1] { // 在后半段
		ans = bs(reverse, len(nums)-1)
	} else if target >= nums[0] && target <= nums[reverse-1] { // 在前半段
		ans = bs(0, reverse-1)
	}

	return ans
}

func search1(nums []int, target int) int {
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

	bs := func(left, right int) int {
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	i := right
	if target > nums[len(nums)-1] { // target 在第一段
		return bs(0, i) // 开区间 (-1, i)
	}
	// target 在第二段
	return bs(i, len(nums)) // 开区间 (i-1, n)
}
