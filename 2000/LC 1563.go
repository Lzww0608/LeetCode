func stoneGameV(a []int) int {
    n := len(a)
    pre := make([]int, n + 1)
    for i, x := range a {
        pre[i+1] = pre[i] + x
    }

    memo := make([][]int, n + 1)
    for i := range memo {
        memo[i] = make([]int, n + 1)
    }

    var dfs func(int, int) int
    dfs = func(l, r int) (res int) {
        if l == r {
            return 
        } else if memo[l][r] != 0 {
            return memo[l][r]
        } 
        defer func() {memo[l][r] = res} ()

        for k := l; k < r; k++ {
            a, b := pre[r] - pre[k], pre[k] - pre[l-1]
            if a > b {
                res = max(res, b + dfs(l, k))
            } else if a < b {
                res = max(res, a + dfs(k + 1, r))
            } else {
                res = max(res, b + dfs(l, k), a + dfs(k + 1, r))
            }
        }
        return
    }
    return dfs(1, n)
}