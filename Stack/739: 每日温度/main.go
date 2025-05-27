package main

func main() {

}

func dailyTemperatures(temperatures []int) []int {
	minStack := make([]int, 0) // 单调栈，只存储最小的索引
	ans := make([]int, len(temperatures))

	for i, t := range temperatures {
		for len(minStack) > 0 && t > temperatures[minStack[len(minStack)-1]] {
			top := minStack[len(minStack)-1] // top 为正在处理的温度较小一天的索引
			ans[top] = i - top               // 更新这一天的答案
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, i)
	}

	return ans
}
