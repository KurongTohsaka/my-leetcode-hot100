package main

import "slices"

func main() {

}

func maxProduct(nums []int) int {
	n := len(nums)
	// nums 中存在负数，所以需要分别存储最大正数乘积和最小负数乘积
	fMax := make([]int, n) // 最大乘积
	fMin := make([]int, n) // 最小乘积
	fMax[0], fMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		// 把 x 加到右端点为 i-1 的（乘积最大/最小）子数组后面，
		// 或者单独组成一个子数组，只有 x 一个元素
		fMax[i] = max(fMax[i-1]*x, fMin[i-1]*x, x)
		fMin[i] = min(fMax[i-1]*x, fMin[i-1]*x, x)
	}
	return slices.Max(fMax)
}
