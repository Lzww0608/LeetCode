const MOD int = 1_000_000_007
func numberOfStableArrays(zero int, one int, limit int) int {
    memo := make([][][2]int, zero + 1)
    for i := range memo {
        memo[i] = make([][2]int, one + 1)
        for j := range memo[i] {
            memo[i][j] = [2]int{-1, -1}
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, j, k int) (res int) {
        if i == 0 {
            if k == 1 && j <= limit {
                return 1
            }
            return
        }
        if j == 0 {
            if k == 0 && i <= limit {
                return 1
            }
            return
        }

        p := &memo[i][j][k]
        if *p != -1 {
            return *p
        }
        defer func() {*p = res} ()

        if k == 0 {
            res = (res + dfs(i - 1, j, 0) + dfs(i - 1, j, 1)) % MOD
            if i > limit {
                res = (res + MOD - dfs(i - limit - 1, j, 1)) % MOD
            }
        } else {
            res = (res + dfs(i, j - 1, 0) + dfs(i, j - 1, 1)) % MOD
            if j > limit {
                res = (res + MOD - dfs(i, j - limit - 1, 0)) % MOD
            }
        }

        return
    }


    return (dfs(zero, one, 0) + dfs(zero, one, 1)) % MOD
}