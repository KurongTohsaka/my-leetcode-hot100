package main

func main() {

}

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	i := 0

	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			// 关键点：不增加i，因为交换回的新元素需要再次检查
		} else {
			i++ // 元素为1时直接跳过
		}
	}
}
