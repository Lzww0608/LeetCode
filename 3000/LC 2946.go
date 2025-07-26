func areSimilar(mat [][]int, k int) bool {
    m, n := len(mat), len(mat[0])
    k %= n 
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][(j + k + n) % n] != mat[i][j] {
                return false
            }
        }
    }

    return true
}