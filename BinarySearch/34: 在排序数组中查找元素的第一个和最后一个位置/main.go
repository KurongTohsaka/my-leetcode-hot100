package main

func main() {

}

func searchRange(nums []int, target int) []int {
	ans := make([]int, 0)

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		// 先找到 target 所在的序列
		if nums[mid] == target {
			i, j := mid, mid
			// 往前找
			for i >= 0 && nums[i] == target {
				i--
			}
			// 往后找
			for j <= len(nums)-1 && nums[j] == target {
				j++
			}
			ans = append(ans, i+1)
			ans = append(ans, j-1)
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if len(ans) == 0 {
		return []int{-1, -1}
	}

	return ans
}
