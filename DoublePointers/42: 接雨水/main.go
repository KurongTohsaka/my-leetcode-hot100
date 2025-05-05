package main

func main() {

}

func trap(height []int) int {
	// preMax、sufMax 分别代表 left 到数组首部的最大高度、right 到数组尾部的最大高度
	left, right, preMax, sufMax := 0, len(height)-1, 0, 0
	ans := 0

	for left < right {
		// 时刻保持更新
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])

		// 木桶效应，盛水容量由短板决定
		// 每个位置的接水量实际上是：min(leftMax[i], rightMax[i]) - height[i]（有了这一步好理解多了）
		if preMax < sufMax {
			// preMax 小，用它和 left 来计算容量
			ans += preMax - height[left]
			left++
		} else {
			// sufMax 小，用它和 right 来计算容量
			ans += sufMax - height[right]
			right--
		}
	}

	return ans
}
