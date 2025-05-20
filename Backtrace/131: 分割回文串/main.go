package main

import "slices"

func main() {

}

func partition(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)
	n := len(s)

	// 辅助函数：判断字符串是否为回文串
	isPalindrome := func(sub string) bool {
		if len(sub) == 0 {
			return false
		}
		l, r := 0, len(sub)-1
		for l < r {
			if sub[l] != sub[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if startIndex == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		for i := startIndex; i < n; i++ {
			substring := s[startIndex : i+1]
			if isPalindrome(substring) {
				path = append(path, substring)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return ans
}
