func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    if m == 0 || n == 0 {
        return
    }

    row := make([]int, m)
    col := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                row[i], col[j] = 1, 1
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if row[i] == 1 || col[j] == 1 {
                matrix[i][j] = 0
            }
        }
    }

    return 
}