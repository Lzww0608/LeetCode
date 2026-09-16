func cyclicShift(n int, g [][]int, rowShift []int, colShift []int) [][]int {
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    for i, x := range rowShift {
        for j := range n {
            ans[i][j] = g[i][(j + x) % n]
        }
    }

    for j, x := range colShift {
        for i := range n {
            g[i][j] = ans[(i + x) % n][j]
        }
    }

    return g
}