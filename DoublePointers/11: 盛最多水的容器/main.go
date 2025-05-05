package main

func main() {

}

func maxArea(height []int) int {
	// 经典双指针
	left, right := 0, len(height)-1
	ans := 0

	for left < right {
		tmp := (right - left) * min(height[left], height[right])
		ans = max(ans, tmp)
		// 哪个指针小移动哪个
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}
