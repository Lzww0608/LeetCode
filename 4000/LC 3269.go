const inf int = math.MaxInt32
func minLargest(a []int, b []int) int {
    calc := func(x, y int) int {
        return x + 2 - ((x & 1) ^ y)
    }
    n, m := len(a), len(b)
    memo := make([][]int, n + 1) 
    for i := range memo {
        memo[i] = make([]int, m + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        p := &memo[i][j]
        if *p != -1 {
            return *p
        }
        if i == 0 && j == 0 {
            return 0
        }
        x, y := inf, inf
        if i > 0 {
            x = calc(dfs(i - 1, j), a[i-1])
        }
        if j > 0 {
            y = calc(dfs(i, j - 1), b[j-1])
        }
        *p = min(x, y)
        return *p
    }
    return dfs(len(a), len(b))
}