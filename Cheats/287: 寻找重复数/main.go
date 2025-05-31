package main

func main() {

}

func findDuplicate(nums []int) int {
	n := len(nums)
	// 遇到数 x ，就把索引为 x 的数加 n
	for i := range nums {
		index := nums[i]%n - 1
		nums[index] += n
	}

	// 当某个数的大于 2*n 即为所求（被连续标记了两次）
	for i := range nums {
		if nums[i] > 2*n {
			return i + 1
		}
	}

	return -1
}
