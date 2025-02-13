func maximumAmount(coins [][]int) int {
    m, n := len(coins), len(coins[0])
    f := make([][][3]int, m + 1)
    for i := range f {
        f[i] = make([][3]int, n + 1)
        for j := range f[i] {
            f[i][j] = [3]int{math.MinInt32 / 2, math.MinInt32 / 2, math.MinInt32 / 2}
        }
    }
    f[1][1][0] = coins[0][0]
    f[1][1][1] = 0

    for i := 0; i < m; i++ {
        for j, x := range coins[i] {
            if i == 0 && j == 0 {
                continue
            }
            f[i+1][j+1][0] = max(f[i+1][j][0], f[i][j+1][0]) + x
            f[i+1][j+1][1] = max(max(f[i+1][j][1], f[i][j+1][1]) + x, f[i+1][j][0], f[i][j+1][0]);
            f[i+1][j+1][2] = max(max(f[i+1][j][2], f[i][j+1][2]) + x, f[i+1][j][1], f[i][j+1][1]);
        }
    }

    return max(f[m][n][0], f[m][n][1], f[m][n][2])
}