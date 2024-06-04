func countSquares(matrix [][]int) int {
    ans := 0
    dp := make([][]int, len(matrix))
    for i := range dp {
        dp[i] = make([]int, len(matrix[0]))
    }

    for i := range matrix {
        for j := range matrix[i] {
            if i == 0 || j == 0 {
                dp[i][j] = matrix[i][j] 
                if dp[i][j] == 1 {
                    ans++
                } 
            } else if matrix[i][j] == 1 {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
                ans += dp[i][j]
            }
        }
    }

    return ans
}
