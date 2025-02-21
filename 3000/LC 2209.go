func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
    n := len(floor)
    memo := make([][]int, numCarpets + 1)
    for i := range memo {
        memo[i] = make([]int, n + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if j <= i * carpetLen {
            return 0
        } else if i == 0 {
            return strings.Count(floor[:j], "1")
        }

        p := &memo[i][j];
        if *p != -1 {
            return *p
        }

        res := dfs(i - 1, j - carpetLen)
        if floor[j-1] == '1' {
            res = min(res, dfs(i, j - 1) + 1)
        } else {
            res = min(res, dfs(i, j - 1))
        }
        *p = res
        return res
    }

    return dfs(numCarpets, n)
}