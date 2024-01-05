var dir [][]int = [][]int{[]int{2,1},[]int{2,-1},[]int{1,2},[]int{1,-2},[]int{-2,1},[]int{-2,-1},[]int{-1,2},[]int{-1,-2}}
func knightProbability(n int, k int, row int, column int) float64 {
    dp := make([][][]float64, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]float64, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]float64, k+1)
            dp[i][j][0] = 1
        }
    }
    for p := 1; p <= k; p++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                for _, d := range dir {
                    x, y := i + d[0], j + d[1]
                    if x < 0 || x >= n || y < 0 || y >= n {
                        continue
                    }
                    dp[i][j][p] += dp[x][y][p-1] / 8
                }
            }
        }
    }
    return dp[row][column][k]

}