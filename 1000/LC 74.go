func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])

    l, r := 0, n * m
    for l < r {
        mid := l + ((r - l) >> 1)
        if matrix[mid / n][mid % n] > target {
            r = mid 
        } else if matrix[mid / n][mid % n] < target {
            l = mid + 1
        } else {
            return true
        }
    }

    return false
}
