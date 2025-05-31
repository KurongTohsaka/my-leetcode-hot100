package main

func main() {

}

func climbStairs(n int) int {
	cache := make([]int, n+1) // 登上第 i 阶梯的方法总数
	cache[0] = 1
	cache[1] = 1

	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}
