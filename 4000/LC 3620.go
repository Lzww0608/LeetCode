func findMaxPathScore(edges [][]int, online []bool, k int64) int {
    n := len(online)
    type pair struct {
        u, c int
    }
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, cost := edge[0], edge[1], edge[2]
        if online[u] && online[v] {
            g[u] = append(g[u], pair{v, cost})
        } 
    }

    memo := make([]int, n)
    check := func(mid int) bool {
        for i := range memo {
            memo[i] = -1
        }
        var dfs func(int) int 
        dfs = func(x int) int {
            if x == n - 1 {
                return 0
            } 
            if memo[x] != -1 {
                return memo[x]
            }
            res := math.MaxInt / 2
            for _, v := range g[x] {
                if v.c >= mid {
                    res = min(res, dfs(v.u) + v.c)
                }
            }
            memo[x] = res
            return res
        }

        return dfs(0) <= int(k)
    }

    l, r := -1, int(k) + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid
        } else {
            r = mid
        }
    }

    return l
}