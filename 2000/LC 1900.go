func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
    type pair struct {mn, mx int}
    f := make([][][]pair, n + 1)
    for i := range f {
        f[i] = make([][]pair, n)
        for j := range f[i] {
            f[i][j] = make([]pair, n)
        }
    }

    var dfs func(int, int, int) pair
    dfs = func(n, first, second int) (ans pair) {
        if first + second == n - 1 {
            return pair{1, 1}
        }

        if first >= n - 1 - first || first > n - 1 - second {
            first, second = n - 1 - second, n - 1 - first
        }

        p := &f[n][first][second]
        if p.mn > 0 {
            return *p
        }

        ans.mn = int(1e9)
        mid := (n + 1) / 2
        var next pair
        for i := 0; i <= first; i++ {
            for j := 0; j < min(second, n - 1 - second) - first; j++ {
                if second < mid {
                    next = dfs(mid, i, i + j + 1)
                } else {
                    next = dfs(mid, i, i + j + 1 + (second * 2 - n + 1) / 2)
                }
                ans.mn = min(ans.mn, next.mn)
                ans.mx = max(ans.mx, next.mx)
            }
        }

        ans.mn++
        ans.mx++
        *p = ans
        return
    }

    res := dfs(n, firstPlayer - 1, secondPlayer - 1)
    return []int{res.mn, res.mx}
}