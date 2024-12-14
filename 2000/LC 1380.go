func luckyNumbers(matrix [][]int) (ans []int) {
    m, n := len(matrix), len(matrix[0])
    col := make(map[int]bool)
    for j := 0; j < n; j++ {
        cur := matrix[0][j]
        for i := 1; i < m; i++ {
            cur = max(cur, matrix[i][j])
        }
        col[cur] = true
    }

    for i := 0; i < m; i++ {
        mn := slices.Min(matrix[i])
        if col[mn] {
            ans = append(ans, mn)
            break
        } 
    }

    return 
}