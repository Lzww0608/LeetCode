func regionsBySlashes(grid []string) int {
    n := len(grid)
    ans := 0
    dp := make([][]int, 3 * n)
    for i := range dp {
        dp[i] = make([]int, 3 * n)
    }
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '/' {
                dp[3*i][3*j+2], dp[3*i+1][3*j+1], dp[3*i+2][3*j] = 1, 1, 1
            } else if grid[i][j] == '\\' {
                dp[3*i][3*j], dp[3*i+1][3*j+1], dp[3*i+2][3*j+2] = 1, 1, 1
            }
        }
    }
    size := 3 * n
    dir := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0},[2]int{-1, 0}}
    var dfs func(x, y int)
    dfs = func(x, y int) {
        if x < 0 || x >= size || y < 0 || y >= size || dp[x][y] == 1{
            return 
        }
        dp[x][y] = 1
        for k := range dir {
            dfs(x + dir[k][0], y + dir[k][1])
        }
    }
    for i := range dp {
        for j := range dp[i] {
            if dp[i][j] == 0 {
                dfs(i, j)
                ans++
            }
        }
    }
    return ans
}
