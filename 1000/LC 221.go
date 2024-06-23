func maximalSquare(matrix [][]byte) (ans int) {
    m, n := len(matrix), len(matrix[0])
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == '1' {
                f[i+1][j+1] = min(f[i][j], f[i+1][j], f[i][j+1]) + 1
                ans = max(ans, f[i+1][j+1] * f[i+1][j+1])
            }
        }
    }
    return
}
