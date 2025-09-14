func numberOfSubmatrices(grid [][]byte) (ans int) {
    type pair struct {
        x, y int
    }

    m, n := len(grid), len(grid[0])
    f := make([]pair, n)
    for i := 0; i < m; i++ {
        tmp := pair{0, 0}
        for j := 0; j < n; j++ {
            if grid[i][j] == 'X' {
                tmp.x++
            } else if grid[i][j] == 'Y' {
                tmp.y++
            }

            f[j].x += tmp.x
            f[j].y += tmp.y 
            if f[j].x == f[j].y && f[j].x != 0 {
                ans++
            }
        }
    }

    return 
}