package main

func main() {

}

func longestConsecutive(nums []int) int {
	// 这里哈希是作为集合使用，保证数字唯一
	hashMap := make(map[int]bool)
	for _, n := range nums {
		hashMap[n] = true
	}
	ans := 0

	for n := range hashMap {
		if hashMap[n-1] {
			continue
		}

		// 不断在集合中寻找连续序列的下一个数字是否存在
		newStart := n + 1
		for hashMap[newStart] {
			newStart++
		}

		ans = max(ans, newStart-n)
	}

	return ans
}
