func countSubTrees(n int, edges [][]int, labels string) []int {
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    ans := make([]int, n)
    var dfs func(int, int) [26]int
    dfs = func(fa, u int) (res [26]int) {
        for _, v := range g[u] {
            if v != fa {
                tmp := dfs(u, v)
                for i, x := range tmp {
                    res[i] += x
                }
            }
        }
        x := int(labels[u] - 'a')
        res[x]++
        ans[u] += res[x]
        return
    }

    dfs(-1, 0)
    return ans
}