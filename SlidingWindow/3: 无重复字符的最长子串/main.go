package main

func main() {}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	left := 0
	// 存储字符及其最新位置（字符是否出现过）
	hashmap := make(map[byte]int)
	ans := 0

	for right, c := range s {
		char := byte(c)
		// 如果字符已存在，更新左指针到 max(left, 上一次出现的位置 + 1)
		if idx, ok := hashmap[char]; ok {
			left = max(left, idx+1)
		}
		// 更新字符的最新位置
		hashmap[char] = right
		ans = max(ans, right-left+1)
	}

	return ans
}
