package main

func main() {

}

// 2 维 DP 解法
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	cache := make([][]bool, n)
	for i := range cache {
		cache[i] = make([]bool, n)
	}
	for i := range n {
		cache[i][i] = true
	}

	maxLen := 1
	begin := 0
	for i := 2; i <= n; i++ {
		for left := 0; left < n; left++ {
			right := i + left - 1
			if right >= n {
				break
			}

			if s[left] != s[right] {
				cache[left][right] = false
			} else {
				if right-left < 3 {
					cache[left][right] = true
				} else {
					cache[left][right] = cache[left+1][right-1]
				}
			}

			if cache[left][right] && right-left+1 > maxLen {
				maxLen = right - left + 1
				begin = left
			}
		}
	}
	return s[begin : begin+maxLen]
}

// 中心扩展解法
func longestPalindrome2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	start, maxLen := 0, 1 // 初始化最长回文子串的起始位置和长度
	for i := 0; i < n; i++ {
		// 检查奇数长度回文（中心为i）
		l1 := expand(s, i, i)
		// 检查偶数长度回文（中心为i和i+1）
		l2 := expand(s, i, i+1)
		// 取两种情况下较长的回文长度
		currentMax := max(l1, l2)
		// 更新最长回文信息
		if currentMax > maxLen {
			maxLen = currentMax
			start = i - (currentMax-1)/2
		}
	}
	return s[start : start+maxLen]
}

// 中心扩展函数：从左右起点向外扩展，返回回文子串长度
func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left-- // 因为是从中心向两侧扩展，所以这里要 left-- 和 right++
		right++
	}
	return right - left - 1
}
