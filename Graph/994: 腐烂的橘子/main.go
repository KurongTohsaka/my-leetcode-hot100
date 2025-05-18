package main

func main() {

}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	fresh := 0                // 新鲜橘子
	queue := make([][]int, 0) // 存储腐烂橘子坐标

	// 统计
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return ans
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// BFS
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			c := queue[0]
			queue = queue[1:]
			for _, d := range directions {
				nextC := []int{c[0] + d[0], c[1] + d[1]}
				// 处理越界和无效坐标
				if nextC[0] < 0 || nextC[0] >= m || nextC[1] < 0 || nextC[1] >= n || grid[nextC[0]][nextC[1]] == 0 {
					continue
				}
				if grid[nextC[0]][nextC[1]] == 1 {
					grid[nextC[0]][nextC[1]] = 2
					fresh--
					queue = append(queue, []int{nextC[0], nextC[1]})
				}

			}
		}
		// 仅当还有腐烂橘子需要处理时才+1
		if len(queue) > 0 {
			ans++
		}
	}

	// 防止有新鲜橘子没有被腐烂
	if fresh > 0 {
		return -1
	}

	return ans
}
