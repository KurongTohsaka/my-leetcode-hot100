package main

func main() {

}

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	dp := make([]int, len(s))
	maxLen := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i) // Push
		} else if len(stack) > 0 && s[i] == ')' {
			leftIndex := stack[len(stack)-1] // 左括号索引
			stack = stack[:len(stack)-1]     // Pop

			length := i - leftIndex + 1 // 当前括号长度
			// 连接前面的有效子串
			if leftIndex-1 >= 0 {
				length += dp[leftIndex-1]
			}
			dp[i] = length
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}
