package main

func main() {

}

func canJump(nums []int) bool {
	maxD := 0 // 最大右侧到达位置

	for i, d := range nums {
		// 不可达
		if i > maxD {
			return false
		}
		maxD = max(maxD, i+d)
	}
	return true
}
