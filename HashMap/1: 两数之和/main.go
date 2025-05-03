package main

func main() {

}

func twoSum(nums []int, target int) []int {
	// 哈希存放元素对应索引
	hashmap := map[int]int{}
	ans := []int{}

	for i, num := range nums {
		if idx, ok := hashmap[target-num]; ok == true {
			ans = append(ans, idx, i)
			return ans
		}
		hashmap[num] = i
	}

	return ans
}
