func checkValid(matrix [][]int) bool {
    n := len(matrix)
    m1 := make(map[int]bool)
    m2 := make(map[int]bool)

    for i := 0; i < n; i++ {
        clear(m1)
        clear(m2)
        for j := 0; j < n; j++ {
            m1[matrix[i][j]] = true
            m2[matrix[j][i]] = true
        }
        if len(m1) != n || len(m2) != n {
            return false
        }
    }

    return true
}