package main

func main() {

}

func rob(nums []int) int {
	n := len(nums)
	cache := make([]int, n+1)
	cache[0], cache[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		cache[i] = max(cache[i-1], cache[i-2]+nums[i-1])
	}

	return cache[n]
}
