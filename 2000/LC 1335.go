func minDifficulty(jobDifficulty []int, d int) int {
    n := len(jobDifficulty)
    if n < d {
        return -1
    }
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, d)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    mx := 0
    for i, x := range jobDifficulty {
        mx = max(mx, x)
        memo[i][0] = mx
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        p := &memo[i][j]
        if *p != -1 {
            return *p
        }

        cur_max := 0
        *p = math.MaxInt32
        for k := i; k >= j; k-- {
            cur_max = max(cur_max, jobDifficulty[k])
            *p = min(*p, cur_max + dfs(k - 1, j - 1))
        }

        return *p
    }

    return dfs(n - 1, d - 1)
}