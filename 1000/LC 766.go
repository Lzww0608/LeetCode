func isToeplitzMatrix(matrix [][]int) bool {
    n, m := len(matrix), len(matrix[0])
    for i := range matrix {
        x, y := i, 0
        for x + 1 < n && y + 1 < m {
            if matrix[x + 1][y + 1] != matrix[x][y] {
                return false
            }
            x, y = x + 1, y + 1
        }
    }

    for j := range matrix[0] {
        x, y := 0, j
        for x + 1 < n && y + 1 < m {
            if matrix[x + 1][y + 1] != matrix[x][y] {
                return false
            }
            x, y = x + 1, y + 1
        }
    }

    return true
}
