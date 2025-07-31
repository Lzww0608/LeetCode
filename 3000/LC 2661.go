func firstCompleteIndex(arr []int, mat [][]int) int {
    m, n := len(mat), len(mat[0])
    pos := make([][2]int, m * n + 1)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            x := mat[i][j]
            pos[x] = [2]int{i, j}
        }
    }

    row := make([]int, m)
    col := make([]int, n)
    for k, x := range arr {
        i, j := pos[x][0], pos[x][1]
        row[i]++
        col[j]++
        if row[i] == n || col[j] == m {
            return k
        }
    }

    return m * n
}