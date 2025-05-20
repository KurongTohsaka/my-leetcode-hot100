package main

func main() {

}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return left
}
