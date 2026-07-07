const MOD int = 1_000_000_007

func countSubsequence(s, t string) int {
    n, m := len(s), len(t)
    if n < m {
        return 0
    }

    f := make([]int, m + 1)
    f[0] = 1
    for i := range n {
        for j := min(i, m - 1); j >= max(m - (n - i), 0); j-- {
            if s[i] == t[j] {
                f[j + 1] += f[j]
            }
        }
    }

    return f[m] % MOD
}

func interleaveCharacters(s1 string, s2 string, target string) int {
    if len(target) > len(s1) + len(s2) {
        return 0
    }

    n, m, d := len(s1), len(s2), len(target)
    memo := make([][][]int, n + 1)
    for i := range memo {
        memo[i] = make([][]int, m + 1)
        for j := range memo[i] {
            memo[i][j] = make([]int, d + 1)
            for k := range memo[i][j] {
                memo[i][j][k] = math.MinInt32
            }
        }
    }

    var dfs func(i, j, k int) int 
    dfs = func(i, j, k int) int {
        if i < -1 || j < -1 || i + j + 1 < k {
            return 0
        } else if k < 0 {
            return 1
        }

        if memo[i + 1][j + 1][k] != math.MinInt32 {
            return memo[i + 1][j + 1][k]
        }

        res := dfs(i, j - 1, k) + dfs(i - 1, j, k) - dfs(i - 1, j - 1, k)
        if i >= 0 && s1[i] == target[k] {
            res += dfs(i - 1, j, k - 1) - dfs(i - 1, j - 1, k - 1)
        }
        if j >= 0 && s2[j] == target[k] {
            res += dfs(i, j - 1, k - 1) - dfs(i - 1, j - 1, k - 1)
        }

        res %= MOD
        memo[i + 1][j + 1][k] = res 
        return res
    }

    ans := dfs(n - 1, m - 1, d - 1) - countSubsequence(s1, target) - countSubsequence(s2, target)
    return (ans % MOD + MOD) % MOD
}