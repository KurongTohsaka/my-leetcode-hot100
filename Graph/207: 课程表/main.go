package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		// 存储所有与 p[1] 课程相关的后续课程
		g[p[1]] = append(g[p[1]], p[0])
	}

	colors := make([]int, numCourses) // 每个节点的颜色
	var dfs func(int) bool
	dfs = func(x int) bool {
		colors[x] = 1 // x 正在访问中
		// colors[y] == 1 ：如果邻居节点 y 的颜色是1，说明在当前的 DFS 路径上又遇到了 y，这就形成了一个环
		// colors[y] == 0 && dfs(y)：递归调用 dfs(y) 继续探索以 y 为起点的路径。如果 dfs(y) 返回 true，表示在以 y 为起点的路径上发现了环，那么以 x 为起点的路径上也存在环
		for _, y := range g[x] {
			if colors[y] == 1 || colors[y] == 0 && dfs(y) {
				return true // 找到了环
			}
		}
		colors[x] = 2 // x 完全访问完毕
		return false  // 没有找到环
	}

	for i, c := range colors {
		if c == 0 && dfs(i) {
			return false // 有环
		}
	}
	return true // 没有环
}
