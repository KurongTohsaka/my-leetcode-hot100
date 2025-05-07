package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	queue := make([]int, 0) // 单调队列，单调递减
	for i, n := range nums {
		// 入，把队列内小于这个数的元素全部出队
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= n {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		// 出，把窗口外的元素出队
		if i-queue[0]+1 > k {
			queue = queue[1:]
		}

		// 记录结果，索引小于窗口时不记录
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}
