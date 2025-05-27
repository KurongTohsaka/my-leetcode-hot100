package main

func main() {

}

func largestRectangleArea(heights []int) (ans int) {
	heights = append(heights, -1) // 最后大火收汁，用 -1 把栈清空
	// 在栈中只有一个数的时候，栈顶的「下面那个数」是 -1，对应 left[i] = -1 的情况
	maxStack := []int{-1} // 最大单调栈
	for right, h := range heights {
		for len(maxStack) > 1 && h <= heights[maxStack[len(maxStack)-1]] {
			i := maxStack[len(maxStack)-1] // 矩形的高（的下标），最大高度
			maxStack = maxStack[:len(maxStack)-1]
			left := maxStack[len(maxStack)-1] // 栈顶下面那个数就是 left，两者相邻
			ans = max(ans, heights[i]*(right-left-1))
		}
		maxStack = append(maxStack, right)
	}
	return
}
