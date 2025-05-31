package main

func main() {

}

func majorityElement(nums []int) int {
	candidate := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			candidate = nums[i]
		}

		// 如果是超过 ⌊1/2⌋ 的值，那么他的票数则永远不会被减为0
		if candidate != nums[i] {
			cnt--
		} else {
			cnt++
		}
	}

	return candidate
}
