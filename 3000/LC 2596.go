func checkValidGrid(grid [][]int) bool {
    n := len(grid)
    type pair struct {
        x, y int
    }
    a := make([]pair, n * n)
    for i := range grid {
        for j, x := range grid[i] {
            a[x] = pair{i, j}
        }
…
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
