func searchMatrix(matrix [][]int, target int) bool {
    n, m := len(matrix), len(matrix[0])

    for i, j := 0, m - 1; i < n && j >= 0; {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] < target {
            i++
        } else {
            j--
        }
    }

    return false
}