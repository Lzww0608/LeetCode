func findBall(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    ans := make([]int, n)
    for j := 0; j < n; j++ {
        x, y := 0, j 
        for x < m {
            t := grid[x][y]
            y += t
            if y < 0 || y >= n || grid[x][y] != t {
                y = -1
                break
            }  
            x++
        }

        ans[j] = y 
    }

    return ans
}