package main

func main() {

}

func jump(nums []int) int {
	// 只维护当前能到达的最右端点和下一次能达到的最右端点
	curRight, nextRight := 0, 0
	ans := 0

	// 只遍历到 n-2，因为 n-1 就已经到目的地了
	for i, num := range nums[:len(nums)-1] {
		// 不断更新下一次最大所能到达的极限
		nextRight = max(nextRight, i+num)
		// i 已经到了当前所能到达的极限，该增长目前的极限了，那就是 nextRight
		if i == curRight {
			curRight = nextRight
			// 每更新一次就相当于跳一次
			ans++
		}
	}

	return ans
}
