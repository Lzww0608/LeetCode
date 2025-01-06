func numWays(n int, relation [][]int, k int) int {
    f := make([][]int, k + 1)
    for i := range f {
        f[i] = make([]int, n)
    }

    f[0][0] = 1
    for i := 0; i < k; i++ {
        for _, v := range relation {
            a, b := v[0], v[1]
            f[i+1][b] += f[i][a]
        }
    }

    return f[k][n-1]
}