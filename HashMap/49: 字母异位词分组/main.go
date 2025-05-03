package main

func main() {}

func groupAnagrams(strs []string) [][]string {
	// 字符计数数组作为键，异位词数组作为值
	hashmap := map[[26]int][]string{}

	for _, str := range strs {
		charCnt := [26]int{}
		for _, char := range str {
			charCnt[char-'a']++
		}
		hashmap[charCnt] = append(hashmap[charCnt], str)
	}

	ans := make([][]string, 0)
	for _, v := range hashmap {
		ans = append(ans, v)
	}

	return ans
}
