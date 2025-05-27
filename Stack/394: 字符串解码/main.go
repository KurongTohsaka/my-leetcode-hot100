package main

import (
	"slices"
	"strconv"
	"unicode"
)

func main() {

}

func decodeString(s string) string {
	stack := make([]rune, 0)

	for _, ch := range s {
		if ch != ']' {
			stack = append(stack, ch)
		} else {
			subStr := make([]rune, 0)
			for stack[len(stack)-1] != '[' {
				subStr = append(subStr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			slices.Reverse(subStr)
			stack = stack[:len(stack)-1] // [ 出栈

			k := make([]rune, 0)
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				k = append(k, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			// 获取重复次数
			slices.Reverse(k)
			cnt, err := strconv.Atoi(string(k))
			if err != nil {
				panic(err)
			}

			// 再次把拼接好的字符串入栈
			for range cnt {
				for _, c := range subStr {
					stack = append(stack, c)
				}
			}
		}
	}

	return string(stack)
}
