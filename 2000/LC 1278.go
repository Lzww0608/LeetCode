func palindromePartition(s string, k int) int {
    n := len(s)
    
    memo := make([][]int, n)
    cost := make([][]int, n)
    for i := range memo {
        cost[i] = make([]int, n)
        memo[i] = make([]int, k)
        for j := range memo[i] {
            memo[i][j] = math.MaxInt32
        }
    }

    for d := 2; d <= n; d++ {
        for i := 0; i <= n - d; i++ {
            j := i + d - 1
            if s[i] == s[j] {
                cost[i][j] = cost[i+1][j-1]
            } else {
                cost[i][j] = cost[i+1][j-1] + 1
            }
        }
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        p := &memo[i][j]
        if *p != math.MaxInt32 {
            return *p
        }

        if j == 0 {
            *p = cost[0][i]
            return *p
        }

        for k := i; k >= j; k-- {
            t := cost[k][i]
            *p = min(*p, t + dfs(k - 1, j - 1))
        }

        return *p
    }

    return dfs(n - 1, k - 1)
}