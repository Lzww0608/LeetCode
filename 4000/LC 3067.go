func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
    n := len(edges) + 1
    ans := make([]int, n)
    type pair struct {
        d, w int
    }
    g := make([][]pair, n)
    for _, e := range edges {
        a, b, c := e[0], e[1], e[2]
        g[a] = append(g[a], pair{b, c})
        g[b] = append(g[b], pair{a, c})
    }

    var dfs func(int, int, int) int
    dfs = func(i, fa, w int) (res int) {
        if w % signalSpeed == 0 {
            res++
        }
        for _, v := range g[i] {
            a, b := v.d, v.w
            if a != fa {
                res += dfs(a, i, w + b)
            }
        }

        return
    }

    for i := range g {
        if len(g[i]) == 1 {
            continue
        }
        sum := 0
        for _, v := range g[i] {
            a, b := v.d, v.w
            cnt := dfs(a, i, b)
            ans[i] += cnt * sum
            sum += cnt
        }
    }

    return ans
}