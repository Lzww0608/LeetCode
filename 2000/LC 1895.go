func largestMagicSquare(a [][]int) int {
    m, n := len(a), len(a[0])
    row := make([][]int, m + 1)
    col := make([][]int, m + 1)
    dg := make([][]int, m + 1)
    udg := make([][]int, m + 1)
    for i := range row {
        row[i] = make([]int, n + 1)
        col[i] = make([]int, n + 1)
        dg[i] = make([]int, n + 1)
        udg[i] = make([]int, n + 1)
    }

    for i := range a {
        for j, x := range a[i] {
            row[i][j+1] = row[i][j] + x 
            col[i+1][j] = col[i][j] +x 
            dg[i+1][j+1] = dg[i][j] + x 
            udg[i+1][j] = udg[i][j+1] + x 
        }
    }

    for k := min(m, n); k > 0; k-- {
    next:
        for i := k; i <= m; i++ {
            for j := k; j <= n; j++ {
                d := dg[i][j] - dg[i-k][j-k]
                if udg[i][j-k] - udg[i-k][j] != d {
                    continue
                }

                for t := i - 1; t >= i - k; t-- {
                    if row[t][j] - row[t][j-k] != d {
                        continue next
                    }
                }

                for t := j - 1; t >= j - k; t-- {
                    if col[i][t] - col[i-k][t] != d {
                        continue next
                    }
                }

                return k
            }
        }
    }

    return 1
}