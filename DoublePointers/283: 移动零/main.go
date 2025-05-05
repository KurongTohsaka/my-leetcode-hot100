package main

func main() {

}

func moveZeroes(nums []int) {
	// 把原始数组当作栈
	stackLen := 0
	for _, num := range nums {
		if num != 0 {
			nums[stackLen] = num
			stackLen++
		}
	}
	clear(nums[stackLen:])
}
