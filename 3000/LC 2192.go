func getAncestors(n int, edges [][]int) [][]int {
    ans := make([][]int, n)
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
    }

    bfs := func(idx int) {
        vis := make([]bool, n)
        q := []int{idx}
        vis[idx] = true
        for len(q) > 0 {
            t := q[0]
            q = q[1:]
            for _, x := range g[t] {
                if !vis[x] {
                    vis[x] =  true
                    q = append(q, x)
                    ans[x] = append(ans[x], idx)
                }
            }
        }
    }

    for i := range ans {
        bfs(i)
    }

 
    return ans
}
