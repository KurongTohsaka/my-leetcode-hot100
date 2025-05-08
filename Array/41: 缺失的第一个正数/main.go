package main

func main() {

}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		// 如果当前学生的学号在 [1,n] 中，但（真身）没有坐在正确的座位上，nums[i] != nums[nums[i]-1] 这个条件是防止重复元素的
		for 1 <= nums[i] && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			// 那么就交换 nums[i] 和 nums[j]，其中 j 是 i 的学号
			j := nums[i] - 1 // 减一是因为数组下标从 0 开始
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 找第一个学号与座位编号不匹配的学生
	for i := range n {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 所有学生都坐在正确的座位上
	return n + 1
}
