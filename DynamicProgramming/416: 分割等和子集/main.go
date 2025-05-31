package main

func main() {}

func canPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	// 无法分为两个子集
	if total%2 != 0 {
		return false
	}

	// 转换为和为一半的 0-1 背包问题
	total /= 2

	// 第一维是指前 i 个元素能否构成子集
	// 第二维指前 i 个元素能否构成和为 j 的子集
	cache := make([][]bool, len(nums)+1)
	for i := range cache {
		cache[i] = make([]bool, total+1)
	}
	cache[0][0] = true

	for i := 0; i < len(nums); i++ {
		for j := 0; j <= total; j++ {
			// 当前元素值大于目标和，跳过
			if j < nums[i] {
				cache[i+1][j] = cache[i][j]
			} else {
				// 选择当前元素时，前 i 个元素能组成 j - nums[i]
				// 不选当前元素时，前 i 个元素能直接组成 j
				cache[i+1][j] = cache[i][j-nums[i]] || cache[i][j]
			}
		}
	}

	return cache[len(nums)][total]
}
