func surfaceArea(grid [][]int) int {
    ans := 0
    n := len(grid)
    for i := range grid {
        for j, c := range grid[i] {
            if c == 0 {
                continue
            }
            ans += 4 * c + 2
            if i + 1 < n {
                ans -= 2 * min(grid[i+1][j], c)
            }  
            if j + 1 < n {
                ans -= 2 * min(grid[i][j+1], c)
            }
        }
    }
    return ans
}
