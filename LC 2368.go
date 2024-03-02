func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
    g := make([][]int, n)
    m := map[int]bool{}
    
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    for _, x := range restricted {
        m[x] = true
    }


    var dfs func(int)
    dfs = func(idx int) {
        if _, ok := m[idx]; ok {
            return
        }
        m[idx] = true
        ans++
        for _, e := range g[idx] {
            dfs(e)
        }
    }
    dfs(0)
    return
}