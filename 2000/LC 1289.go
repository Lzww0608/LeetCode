func minFallingPathSum(grid [][]int) int {
    n := len(grid)
    if n == 1 {
        return grid[0][0]
    }
    
    dp := make([]int, n)
    var x, y int
    for i := range dp {
        dp[i] = grid[0][i]
        if i == 0 || dp[i] < dp[x] {
            y = x
            x = i
        } else if i == 1 || (i != x && dp[i] < dp[y]) {
            y = i
        }
    }


    for i := 1; i < n; i++ {
        next := make([]int, n)
        newx, newy := 0, 0 
        for j := 0; j < n; j++ {
            if j == x {
                next[j] = grid[i][j] + dp[y] 
            } else {
                next[j] = grid[i][j] + dp[x]
            }

            if j == 0 || next[j] < next[newx] {
                newy = newx
                newx = j
            } else if j == 1 || (j != newx && next[j] < next[newy]) {
                newy = j
            }
        }
        dp = next
        x, y = newx, newy 
    }

    return dp[x] 
}
