const MOD int = 1_000_000_007
func subsequencePairCount(nums []int) int {
    m := slices.Max(nums) + 1
    n := len(nums)
    memo := make([][][]int, n)
    for i := 0; i < n; i++ {
        memo[i] = make([][]int, m)
        for j := 0; j < m; j++ {
            memo[i][j] = make([]int, m)
            for k := 0; k < m; k++ {
                memo[i][j][k] = -1
            }
        }
    }

    g := make([][]int, m)
    g[0] = make([]int, m)
    for i := 1; i < m; i++ {
        g[i] = make([]int, m)
        for j := 0; j < m; j++ {
            g[i][j] = gcd(i, j)
        }
    }

    var dfs func(int, int, int) int 
    dfs = func(i, j, k int) int {
        if i < 0 {
            if j == k {
                return 1
            }
            return 0
        }

        p := &memo[i][j][k]
        if *p != -1 {
            return *p
        }

        *p = dfs(i - 1, j, k) + dfs(i - 1, g[nums[i]][j], k) + dfs(i - 1, j, g[nums[i]][k])
        *p %= MOD
        return *p
    }


    return (dfs(n - 1, 0, 0) + MOD - 1) % MOD
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}