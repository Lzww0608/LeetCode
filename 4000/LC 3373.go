func build(edges [][]int) ([][]int, []int) {
    cnt := make([]int, 2)
    n := len(edges) + 1
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    var dfs func(int, int, int)
    dfs = func(u, fa, d int) {
        cnt[d]++
        for _, v := range g[u] {
            if v != fa {
                dfs(v, u, d ^ 1)
            }
        }
    }
    dfs(0, -1, 0)

    return g, cnt
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
    _, cnt2 := build(edges2)
    mx := max(cnt2[0], cnt2[1])

    g, cnt1 := build(edges1)
    n := len(g)
    ans := make([]int, n)
    var dfs func(int, int, int)
    dfs = func(u, fa, d int) {
        ans[u] = mx + cnt1[d]
        for _, v := range g[u] {
            if v != fa {
                dfs(v, u, d ^ 1)
            }
        }
    }
    dfs(0, -1, 0)

    return ans
}