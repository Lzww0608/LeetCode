func specialNodes(n int, edges [][]int, x int, y int, z int) (ans int) {
    g := make([][]int, n)
    for _, v := range edges {
        a, b := v[0], v[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    var dfs func(int, int, int, []int) 
    dfs = func(u, fa, d int, a []int) {
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            a[v] = d
            dfs(v, u, d + 1, a)
        }
    }

    a := [3][]int{}
    b := [3]int{x, y, z}
    for i := range a {
        a[i] = make([]int, n)
        dfs(b[i], -1, 1, a[i])
    }
    
    for i := range n {
        c := []int{a[0][i], a[1][i], a[2][i]}
        sort.Ints(c)
        if c[0] * c[0] + c[1] * c[1] == c[2] * c[2] {
            ans++
        }
    }

    return
}