const MOD int = 1_000_000_007
func countWinningSequences(s string) int {
    n := len(s)
    m := map[byte]int {
        'F': 0,
        'W': 1,
        'E': 2,
    }

    pow := make([]int, n + 1)
    pow[0] = 1
    for i := 1; i <= n; i++ {
        pow[i] = pow[i - 1] * 2 % MOD
    }

    memo := make([][][3]int, n)
    for i := range memo {
        memo[i] = make([][3]int, n * 2 + 1)
        for j := range memo[i] {
            memo[i][j] = [3]int{-1, -1, -1}
        }
    }

    var dfs func(int, int, int) int 
    dfs = func(i, j, pre int) (res int) {
        if i + j + 1 <= 0 {
            return 0
        }
        if j > i + 1 {
            return pow[i + 1]
        }

        p := &memo[i][j + n][pre]
        if *p != -1 {
            return *p
        }
        for k := 0; k < 3; k++ {
            if i == n - 1 || k != pre {
                score := (k - m[s[i]] + 3) % 3
                if score == 2 {
                    score = -1
                }
                res = (res + dfs(i - 1, j + score, k)) % MOD
            }
        }
        *p = res
        return
    }

    return dfs(n - 1, 0, 0)
}