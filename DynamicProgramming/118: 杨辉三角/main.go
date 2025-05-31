package main

func main() {

}

func generate(numRows int) [][]int {
	cache := make([][]int, numRows)

	// 每行的首尾都是 1
	for i := 1; i < numRows+1; i++ {
		cache[i-1] = make([]int, i)
		cache[i-1][0], cache[i-1][i-1] = 1, 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			cache[i][j] = cache[i-1][j-1] + cache[i-1][j]
		}
	}

	return cache
}
