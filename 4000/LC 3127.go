func canMakeSquare(grid [][]byte) bool {
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            cnt := 0
            if grid[i][j] == 'B' {
                cnt++
            } else {
                cnt--
            }
            if grid[i+1][j] == 'B' {
                cnt++
            } else {
                cnt--
            }
            if grid[i][j+1] == 'B' {
                cnt++
            } else {
                cnt--
            }
            if grid[i+1][j+1] == 'B' {
                cnt++
            } else {
                cnt--
            }

            if cnt != 0 {
                return true
            }
        }
    }

    return false
}