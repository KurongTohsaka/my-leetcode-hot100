package main

func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				cache[i][j] = max(cache[i-1][j], cache[i][j-1])
			}
		}
	}

	return cache[m][n]
}
