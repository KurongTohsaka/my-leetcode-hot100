package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	// 统计每个元素的出现次数
	cnt := map[int]int{}
	maxCnt := 0
	for _, x := range nums {
		cnt[x]++
		maxCnt = max(maxCnt, cnt[x])
	}

	// 桶排序，根据统计次数存储对应元素
	buckets := make([][]int, maxCnt+1)
	for i, c := range cnt {
		buckets[c] = append(buckets[c], i)
	}

	// 倒序遍历桶
	ans := make([]int, 0, k)
	for i := len(buckets); i >= 0 && len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}

	return ans
}
