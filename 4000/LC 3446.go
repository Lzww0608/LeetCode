func sortMatrix(grid [][]int) [][]int {
    n := len(grid)

    for i := 0; i < n - 1; i++ {
        a := make([]int, 0, n - i) 
        for j := 0; i + j < n; j++ {
            a = append(a, grid[i + j][j])
        }

        sort.Slice(a, func(i, j int) bool {
            return a[i] > a[j]
        })

        for j := 0; i + j < n; j++ {
            grid[i + j][j] = a[j]
        }
    }

    for j := 1; j < n - 1; j++ {
        a := make([]int, 0, n - j) 
        for i := 0; i + j < n; i++ {
            a = append(a, grid[i][i + j])
        }

        sort.Ints(a)

        for i := 0; i + j < n; i++ {
            grid[i][i + j] = a[i]
        }
    }

    return grid
}