func colorBorder(grid [][]int, row int, col int, color int) [][]int {
    m, n := len(grid), len(grid[0])
    q := []int{}
    c := grid[row][col]
    vis := make([]bool, m * n)
    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != c {
            return false
        }
        if vis[i * n + j] {
            return true
        }
        vis[i * n + j] = true
        border := dfs(i + 1, j) 
        border = dfs(i, j + 1) && border
        border = dfs(i - 1, j) && border
        border = dfs(i, j - 1) && border 
        if !border {
            q = append(q, i * n + j)
        }
        return true
    }
    dfs(row, col)

    for _, v := range q {
        grid[v / n][v % n] = color
    }

    return grid 
}