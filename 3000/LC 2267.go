func hasValidPath(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    memo := make([][][]int, m)
    for i := range memo {
        memo[i] = make([][]int, n)
        for j := range memo[i] {
            memo[i][j] = make([]int, 400)
            for k := range memo[i][j] {
                memo[i][j][k] = -1
            }
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, j, k int) int {
        if i >= m || j >= n {
            return 0
        }
        if grid[i][j] == '(' {
            k++
        } else {
            k--
        }

        if i == m - 1 && j == n - 1 {
            if k == 0 {
                return 1
            } else {
                return 0
            }
        }

        p := &memo[i][j][k+200]
        if *p != -1 || k < 0 {
            *p = -1
            return *p
        }

        if dfs(i + 1, j, k) == 1 {
            *p = 1
            return 1
        } 
        if dfs(i, j + 1, k) == 1 {
            *p = 1
            return 1
        }
        *p = 0
        return 0
    }

    if dfs(0, 0, 0) == 1 {
        return true
    }

    return false
}