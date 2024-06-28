func modifiedMatrix(matrix [][]int) [][]int {
    n := len(matrix[0])
    a := make([]int, n)
    for i := range a {
        a[i] = -10
    }

    for j := 0; j < n; j++ {
        for i := range matrix {
            a[j] = max(a[j], matrix[i][j])
        }  
    }

    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == -1 {
                matrix[i][j] = a[j]
            }
        }
    }

    return matrix
}
