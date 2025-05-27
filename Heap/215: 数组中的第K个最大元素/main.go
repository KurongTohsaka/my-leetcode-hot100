package main

func main() {

}

func findKthLargest(nums []int, k int) int {
	size := len(nums)
	buildMaxHeap(nums, size)

	for i := 0; i < k-1; i++ {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		maxHeapify(nums, 0, size)
	}

	return nums[0]
}

// buildMaxHeap 从底向上构建最大堆
func buildMaxHeap(nums []int, size int) {
	for i := size/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, size)
	}
}

// maxHeapify 调整堆
func maxHeapify(nums []int, i, size int) {
	left := i*2 + 1
	right := size*2 + 1
	maxIdx := i

	if left < size && nums[left] > nums[maxIdx] {
		maxIdx = left
	}
	if right < size && nums[right] > nums[maxIdx] {
		maxIdx = right
	}
	if maxIdx != i {
		nums[maxIdx], nums[i] = nums[i], nums[maxIdx]
		// 继续调整以 maxIdx 为根的子树
		maxHeapify(nums, maxIdx, size)
	}
}
