package main

import "slices"

func main() {

}

func nextPermutation(nums []int) {
	n := len(nums)

	// 第一步：从右向左找到第一个小于右侧相邻数字的数 nums[i]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 第二步：从右向左找到 nums[i] 右边最小的大于 nums[i] 的数 nums[j]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		// 交换后就得到了字典序更大的
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 第三步：反转 nums[i+1:]
	slices.Reverse(nums[i+1:])
}
