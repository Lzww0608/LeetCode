const MOD int = 1_000_000_007
func numOfArrays(n int, m int, k int) (ans int) {
    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, m + 1)
        for j := range memo[i] {
            memo[i][j] = make([]int, k)
            for t := range memo[i][j] {
                memo[i][j][t] = -1
            }
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, j, t int) (res int) {
        if t > n - i || t < 0 {
            return 0
        } else if i == n {
            return 1
        }

        if memo[i][j][t] != -1 {
            return memo[i][j][t]
        }
        
        res = dfs(i + 1, j, t) * j % MOD
        for d := j + 1; d <= m; d++ {
            res = (res + dfs(i + 1, d, t - 1)) % MOD
        }

        memo[i][j][t] = res
        return
    }

    for i := 1; i <= m; i++ {
        ans = (ans + dfs(1, i, k - 1)) % MOD
    }

    return
}