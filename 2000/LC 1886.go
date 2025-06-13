func findRotation(mat [][]int, target [][]int) bool {
    n := len(mat)
    a, b, c, d := true, true, true, true
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if a && mat[i][j] != target[n-j-1][i] {
                a = false
            }

            if b && mat[i][j] != target[n-i-1][n-1-j] {
                b = false
            }

            if c && mat[i][j] != target[j][n-i-1] {
                c = false
            }

            if d && mat[i][j] != target[i][j] {
                d = false
            }
        }
    }

    return a || b || c || d
}