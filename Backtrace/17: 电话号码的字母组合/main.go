package main

import "slices"

func main() {

}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans := make([]string, 0)
	path := make([]byte, 0)
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(slices.Clone(path)))
			return
		}

		for _, c := range phoneMap[string(digits[i])] {
			path = append(path, byte(c))
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return ans
}
