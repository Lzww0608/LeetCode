func calculateMinimumHP(dungeon [][]int) int {
    m, n := len(dungeon), len(dungeon[0])
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }

    f[m][n-1], f[m-1][n] = 1, 1
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            cur := min(f[i+1][j], f[i][j+1]) - dungeon[i][j]
            if cur <= 0 {
                f[i][j] = 1
            } else {
                f[i][j] = cur
            }
        }
    }

    return f[0][0]
}