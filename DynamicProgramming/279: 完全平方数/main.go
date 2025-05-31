package main

func main() {

}

func numSquares(n int) int {
	cache := make([]int, n+1) // 每个和为 i 的最少完全平方数
	cache[0] = 0
	for i := 1; i <= n; i++ {
		cache[i] = i
		// 枚举
		for j := 1; j*j <= i; j++ {
			cache[i] = min(cache[i], cache[i-j*j]+1)
		}
	}
	return cache[n]
}
