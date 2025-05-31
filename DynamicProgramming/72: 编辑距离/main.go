package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for j := range n {
		// 将空字符串转换为 word2 的前 j 个字符，需要进行 j 次插入操作
		cache[0][j+1] = j + 1
	}

	for i := 1; i <= m; i++ {
		// 将 word1 的前 i 个字符转换为空字符串，需要进行 i 次删除操作
		cache[i+1][0] = i + 1
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				cache[i][j] = cache[i-1][j-1]
			} else {
				// 三种操作取最小值
				cache[i][j] = min(cache[i-1][j], cache[i][j-1], cache[i-1][j-1]) + 1
			}
		}
	}

	return cache[m][n]
}
