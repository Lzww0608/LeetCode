
const N int = 20

func interactionCosts(n int, edges [][]int, group []int) int64 {
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	type pair struct {
		x, y int
		// x: 子树中，与当前节点同组的节点数量
		// y: 子树中，与当前节点同组的节点距离当前节点的距离之和
	}
	ans := 0
	var dfs func(int, int) (cnt [N]pair)
	dfs = func(u, fa int) (cnt [N]pair) {
		x := group[u] - 1
		cnt[x].x++
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			cur := dfs(v, u)
			for i := range N {
				if cnt[i].x > 0 && cur[i].x > 0 {
					ans += cur[i].x*(cnt[i].y) + cnt[i].x*(cur[i].y+cur[i].x)
				}
			}
			for i := range N {
				cnt[i].x += cur[i].x
				cnt[i].y += cur[i].y + cur[i].x
			}
		}

		return
	}
	dfs(0, -1)

	return int64(ans)
}
