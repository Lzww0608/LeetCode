func findSquare(matrix [][]int) (ans []int) {
    n := len(matrix)
    if n == 0 {
        return
    }

    left := make([][]int, n + 1)
    up := make([][]int, n + 1)
    for i := range left {
        left[i] = make([]int, n + 1)
        up[i] = make([]int, n + 1)
    }

    mx := 0
    for i := range matrix {
        for j, x := range matrix[i] {
            if x == 1 {
                continue
            }
            left[i+1][j+1] = left[i+1][j] + 1
            up[i+1][j+1] = up[i][j+1] + 1
            d := min(left[i+1][j+1], up[i+1][j+1])
            for d > 0 && (left[i+2-d][j+1] < d || up[i+1][j+2-d] < d) {
                d--
            }
            if d > mx {
                mx = d
                ans = []int{i + 1 - d, j + 1 - d, d}
            }
        }
    }

    return
}