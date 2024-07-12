const MOD int = 1_000_000_007
func dieSimulator(n int, rollMax []int) (ans int) {
    memo := make([][6][15]int, n)
    for i := range memo {
        for j := range memo[i] {
            for k := range memo[i][j] {
                memo[i][j][k] = -1
            }
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, last, left int) int {
        if i == 0 {
            return 1
        }
        p := &memo[i][last][left]
        if *p >= 0 {
            return *p
        }
        ans := 0
        for j := 0; j < 6; j++ {
            if j != last {
                ans += dfs(i - 1, j, rollMax[j] - 1)
            } else if left > 0 {
                ans += dfs(i - 1, j, left - 1)
            }
        }
        *p = ans % MOD
        return *p
    }

    for i := 0; i < 6; i++ {
        ans += dfs(n - 1, i, rollMax[i] - 1)
    }
    return ans % MOD
}