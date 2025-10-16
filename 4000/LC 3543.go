func maxWeight(n int, edges [][]int, k int, t int) (ans int) {
    if k > n {
        return -1
    }

    type pair struct {
        u, w int 
    }

    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w})
    }

    memo := make(map[int]bool)
    ans = -1

    var dfs func(int, int, int)
    dfs = func(x, cnt, sum int) {
        if cnt == k {
            ans = max(ans, sum)
            return
        }
        
        mask := (x << 20) | (cnt << 10) | sum 
        if _, ok := memo[mask]; ok {
            return
        }
        memo[mask] = true

        for _, v := range g[x] {
            if sum + v.w < t {
                dfs(v.u, cnt + 1, sum + v.w)
            }
        }

        return
    }
    
    for i := range n {
        dfs(i, 0, 0)
    }

    return
}