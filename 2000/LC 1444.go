const MOD int = 1_000_000_007
func ways(pizza []string, k int) int {
    m, n := len(pizza), len(pizza[0])
    sum := make([][]int, m + 1)
    for i := range sum {
        sum[i] = make([]int, n + 1)
    }

    for i, row := range pizza {
        for j, x := range row {
            sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + int(x & 1)
        }
    }

    query := func(r1, c1, r2, c2 int) bool {
        return sum[r2][c2] - sum[r1][c2] - sum[r2][c1] + sum[r1][c1] > 0
    }

    memo := make([][][]int, k)
    for i := range memo {
        memo[i] = make([][]int, m)
        for j := range memo[i] {
            memo[i][j] = make([]int, n)
            for t := range memo[i][j] {
                memo[i][j][t] = -1
            }
        }
    }

    var dfs func(int, int, int) int
    dfs = func(c, i, j int) (res int) {
        if c == 0 {
            if query(i, j, m, n) {
                return 1
            }
            return
        }

        p := &memo[c][i][j]
        if *p != -1 {
            return *p
        }
        defer func() {*p = res} ()

        for i2 := i + 1; i2 < m; i2++ {
            if query(i, j, i2, n) {
                res = (res + dfs(c - 1, i2, j)) % MOD
            }
        }
        for j2 := j + 1; j2 < n; j2++ {
            if query(i, j, m, j2) {
                res = (res + dfs(c - 1, i, j2)) % MOD
            }
        }

        return
    }

    return dfs(k - 1, 0, 0)
}