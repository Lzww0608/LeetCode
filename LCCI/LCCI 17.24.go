func getMaxMatrix(a [][]int) (ans []int) {
    m, n := len(a), len(a[0])
    ans = []int{0, 0, 0, 0}

    pre := make([][]int, m + 1)
    for i := range pre {
        pre[i] = make([]int, n + 1)
    }

    for i := range a {
        for j, x := range a[i] {
            pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + x
        }
    }

    mx := math.MinInt
    for i := range a {
        for j := i; j < m; j++ {
            cur, l := 0, 0
            for r := 0; r < n; r++ {
                cur = pre[j+1][r+1] + pre[i][l] - pre[i][r+1] - pre[j+1][l]
                if mx < cur {
                    mx = cur
                    ans = []int{i, l, j, r}
                }
                if cur < 0 {
                    cur = 0
                    l = r + 1
                }
            }
        }
    }

    return
}