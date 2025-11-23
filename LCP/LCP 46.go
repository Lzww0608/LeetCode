func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	ans := make([]int, n)
	copy(ans[1:], finalCnt)

	p := make([]int, n)
	p[0] = 1
	g := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	for i := len(plans) - 1; i >= 0; i-- {
		num, idx := plans[i][0], plans[i][1]
		if num == 1 {
			ans[idx] <<= 1
			p[idx] <<= 1
		} else if num == 2 {
			for _, j := range g[idx] {
				ans[j] -= ans[idx]
				p[j] -= p[idx]
			}
		} else {
			for _, j := range g[idx] {
				ans[j] += ans[idx]
				p[j] += p[idx]
			}
		}
	}

	a, b := 0, 0
	for i := range n {
		a += ans[i]
		b += p[i]
	}

	rate := (int(totalNum) - a) / b
	for i := range n {
		ans[i] += p[i] * rate
	}

	return ans
}