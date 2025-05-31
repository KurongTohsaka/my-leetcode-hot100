package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		words[w] = true
	}

	n := len(s)
	cache := make([]bool, n+1) // 到第 i 为止的子串能否被表示
	cache[0] = true
	for i := 1; i <= n; i++ {
		// 假设单词 leetcode，leet 已经存在，那只需去寻找到 i-1 为止的子串能否被 code 及其子串表示
		// 然后不断递归子问题
		for j := 0; j < i; j++ {
			// 当前索引 j 为止的子串可以被表示且 [j, i-1] 的子串存在于词汇表 words 中
			// i=4：检查 j=0，s[0:4]="leet" 在字典中，且 cache[0]=true → cache[4]=true
			// i=8：检查 j=4，s[4:8]="code" 在字典中，且 cache[4]=true → cache[8]=true
			if cache[j] && words[s[j:i]] {
				cache[i] = true
				break
			}
		}
	}

	return cache[n]
}
