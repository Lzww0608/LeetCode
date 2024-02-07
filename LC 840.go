func numMagicSquaresInside(grid [][]int) int {
    ans := 0
    n, m := len(grid), len(grid[0])

    check := func(x, y int) bool {
        if grid[x+1][y+1] != 5 {
            return false
        }
        m := map[int]int{}
        for i := x; i < x + 3; i++ {
            for j := y; j < y + 3; j++ {
                m[grid[i][j]]++
            }
        }
        for i := 1; i < 10; i++ {
            if m[i] != 1 {
                return false
            }
        }
        sum := grid[x][y] + grid[x][y+1] + grid[x][y+2]
        for i := x + 1; i < x + 3; i++ {
            sum_x := 0
            for j := y; j < y + 3; j++ {
                sum_x += grid[i][j]
            }
            if sum != sum_x {
                return false
            }
        }

        for j := y; j < y + 3; j++ {
            sum_y := 0
            for i := x; i < x + 3; i++ {
                sum_y += grid[i][j]
            }
            if sum != sum_y {
                return false
            }
        }

        if grid[x][y] + grid[x+1][y+1] + grid[x+2][y+2] != sum || grid[x][y+2] + grid[x+1][y+1] + grid[x+2][y] != sum {
            return false
        }
        return true
    } 

    for i := 0; i + 3 <= n; i++ {
        for j := 0; j + 3 <= m; j++ {
            if check(i, j) {
                ans++
            }
        }
    } 
    return ans
}