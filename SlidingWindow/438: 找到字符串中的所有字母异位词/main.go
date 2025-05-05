package main

import "slices"

func main() {

}

func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if n < m {
		return nil
	}

	left := 0
	ans := make([]int, 0)
	window := make([]int, 26)
	target := make([]int, 26)
	for _, c := range p {
		target[c-'a']++
	}

	for right, c := range s {
		window[c-'a']++

		// 当窗口大小等于 p 的长度时，开始检查并移动左指针
		if right-left+1 >= m {
			if slices.Equal(window, target) {
				ans = append(ans, left)
			}
			// 移动左指针前，先减少左指针指向的字符计数
			window[s[left]-'a']--
			left++
		}
	}

	return ans
}
