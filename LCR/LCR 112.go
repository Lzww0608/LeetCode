func longestIncreasingPath(matrix [][]int) (ans int) {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    n, m := len(matrix), len(matrix[0])
    rec := make([]int, n * m)

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if rec[i * m + j] > 0 {
            return rec[i * m + j]
        }

        res := 1
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] > matrix[i][j] {
                res = max(res, 1 + dfs(x, y))
            }
        }
        rec[i * m + j] = res
        return res
    }

    for i := range matrix {
        for j := range matrix[i] {
            if rec[i * m + j] == 0 {
                ans = max(ans, dfs(i, j))
            } 
        }
    }

    return 
}