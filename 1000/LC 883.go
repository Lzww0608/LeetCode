func projectionArea(grid [][]int) int {
    area1, area2, area3 := 0, 0, 0
    for i := range grid {
        a, b := 0, 0
        for j, s := range grid[i] {
            if s != 0 {
                area1++
            }
            a = max(a, grid[i][j])
            b = max(b, grid[j][i])
        }
        area2 += a
        area3 += b
    }
    return area1 + area2 + area3
}
