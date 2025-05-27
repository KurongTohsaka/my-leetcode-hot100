package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 计算两个数组的总长度
	total := len(nums1) + len(nums2)

	// 判断总长度是偶数还是奇数
	if total%2 == 0 {
		// 如果总长度是偶数，中位数是中间两个数的平均值。
		// 我们需要找到第 total/2 个数和第 total/2 + 1 个数。
		// 注意：这里的 n 是指第 n 小的数（1-based indexing），所以 findN(..., k) 返回的是第 k 小的数。
		// 因此，对于第 total/2 个数，需要传入 total/2；对于第 total/2+1 个数，需要传入 total/2+1。
		return float64(findN(nums1, nums2, total/2)+findN(nums1, nums2, total/2+1)) / 2
	} else {
		// 如果总长度是奇数，中位数是中间的那个数。
		// 例如，总长度为 5，中位数是第 3 个数。
		// 我们需要找到第 total/2 + 1 个数。
		return float64(findN(nums1, nums2, total/2+1))
	}
}

// findN 递归函数，用于找到两个有序数组合并后的第 n 小的数
// n 表示要找的是第几小的数 (从 1 开始计数)
func findN(nums1 []int, nums2 []int, n int) int {
	// 确保 nums1 是较短的数组，可以简化后续的递归调用和边界处理
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// --- 递归终止条件 ---

	// 1. 如果 nums1 已经为空，说明第 n 小的数一定在 nums2 中
	// 因为 n 是 1-based index，所以对应数组的索引是 n-1
	if len(nums1) == 0 {
		return nums2[n-1]
	}

	// 2. 如果 nums2 已经为空 (且 nums1 不为空，因为上面已确保 nums1 较短)，说明第 n 小的数一定在 nums1 中
	if len(nums2) == 0 {
		return nums1[n-1]
	}

	// 3. 如果 n 等于 1，表示我们要找的是最小的数
	// 直接比较两个数组的第一个元素，返回较小者
	if n == 1 {
		return min(nums1[0], nums2[0])
	}

	// --- 递归步骤 ---

	// 每次递归，我们尝试排除掉 k 个元素（k = n/2）。
	// 在每个数组中都尝试跳过大约 n/2 个元素。
	// index 变量表示我们将在当前数组中考虑跳过的元素数量。
	// 注意：这里 `min` 函数用于防止 `index` 超出数组边界，确保我们不会访问越界的索引。
	index := n / 2
	i1, i2 := min(index, len(nums1)), min(index, len(nums2))

	// 比较 nums1 中第 i1 个元素（即 nums1[i1-1]）和 nums2 中第 i2 个元素（即 nums2[i2-1]）
	// 这里的 i1-1 和 i2-1 是 0-based 索引，对应第 i1 和第 i2 个元素
	if nums1[i1-1] < nums2[i2-1] {
		// 如果 nums1 的第 i1 个元素小于 nums2 的第 i2 个元素：
		// 这意味着 nums1 的前 i1 个元素（nums1[0]...nums1[i1-1]）肯定都比 nums2 的第 i2 个元素小。
		// 所以，nums1 的这 i1 个元素肯定都在我们正在寻找的第 n 小的数之前。
		// 我们可以安全地将 nums1 的前 i1 个元素以及它们之前的部分排除掉。
		// 递归调用 findN：
		//   - nums1 从 i1 之后的子数组开始（nums1[i1:]）
		//   - nums2 保持不变
		//   - 要寻找的数变为第 n-i1 小的数（因为已经排除了 i1 个较小的元素）
		return findN(nums1[i1:], nums2, n-i1)
	} else {
		// 如果 nums1 的第 i1 个元素大于等于 nums2 的第 i2 个元素：
		// 这意味着 nums2 的前 i2 个元素（nums2[0]...nums2[i2-1]）肯定都比 nums1 的第 i1 个元素小。
		// 所以，nums2 的这 i2 个元素肯定都在我们正在寻找的第 n 小的数之前。
		// 我们可以安全地将 nums2 的前 i2 个元素以及它们之前的部分排除掉。
		// 递归调用 findN：
		//   - nums1 保持不变
		//   - nums2 从 i2 之后的子数组开始（nums2[i2:]）
		//   - 要寻找的数变为第 n-i2 小的数（因为已经排除了 i2 个较小的元素）
		return findN(nums1, nums2[i2:], n-i2)
	}
}
