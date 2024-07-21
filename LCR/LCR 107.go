func updateMatrix(mat [][]int) [][]int {
    m, n := len(mat), len(mat[0])

    for i := range mat {
        for j := range mat[i] {
            if mat[i][j] == 1 {
                mat[i][j] = math.MaxInt / 2
            }
            if i > 0 {
                mat[i][j] = min(mat[i][j], mat[i-1][j] + 1)
            }
            if j > 0 {
                mat[i][j] = min(mat[i][j], mat[i][j-1] + 1)
            }
        }
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if i < m - 1 {
                mat[i][j] = min(mat[i][j], mat[i+1][j] + 1)
            }
            if j < n - 1 {
                mat[i][j] = min(mat[i][j], mat[i][j+1] + 1)
            }
        }
    }

    return mat
}