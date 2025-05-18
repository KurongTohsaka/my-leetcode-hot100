package main

func main() {

}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(curNum int)

	dfs = func(curNum int) {
		if curNum == n {
			// 这里一定要对 path 进行深复制，否则直接 append 会导致底层数组没有变化
			ans = append(ans, append([]int(nil), path...))
			return
		} else {
			for i, on := range onPath {
				if !on {
					path[curNum] = nums[i]
					onPath[i] = true
					dfs(curNum + 1)
					onPath[i] = false
				}
			}
		}
	}

	dfs(0)

	return ans
}
