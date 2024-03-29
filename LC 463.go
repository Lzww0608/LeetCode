func islandPerimeter(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])

    //dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    edge := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 {
                continue
            }
            ans++
            if i + 1 < m && grid[i+1][j] == 1 {
                edge++
            }
            if j + 1 < n && grid[i][j+1] == 1 {
                edge++
            }
        }
    }
    ans = ans * 4 - edge * 2
    return
}