func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    m, n := len(matrix), len(matrix[0])
    i, j := 0, n - 1
    for i >= 0 && i < m && j >= 0 && j < n {
        x := matrix[i][j]
        if x == target {
            return true
        } else if x < target {
            i++
        } else {
            j--
        }
    }

    return false
}