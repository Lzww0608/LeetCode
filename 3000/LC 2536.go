func rangeAddQueries(n int, queries [][]int) [][]int {
    dif := make([][]int, n + 2)
    for i := range dif {
        dif[i] = make([]int, n + 2)
    }
    for _, q := range queries {
        x1, y1, x2, y2 := q[0], q[1], q[2], q[3]
        dif[x1+1][y1+1] += 1
        dif[x1+1][y2+2] -= 1
        dif[x2+2][y1+1] -= 1
        dif[x2+2][y2+2] += 1
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            dif[i][j] += dif[i][j-1] + dif[i-1][j] - dif[i-1][j-1]
        }
    }

    dif = dif[1:n+1]
    for i := 0; i < n; i++ {
        dif[i] = dif[i][1:n+1]
    }
    
    return dif
}