package main

func main() {

}

func singleNumber(nums []int) (ans int) {
	for _, x := range nums {
		// 异或，相同的数字的位会被消除，最后只会剩下只出现一次的数字
		ans ^= x
	}
	return
}
