func stoneGameVII(stones []int) int {
    n := len(stones)
    pre := make([]int, n + 1)
    for i, x := range stones {
        pre[i+1] = pre[i] + x
    }

    f := make([][]int, n)
    for i := n - 1; i >= 0; i-- {
        f[i] = make([]int, n)
        for j := i + 1; j < n; j++ {
            f[i][j] = max(pre[j+1] - pre[i+1] - f[i+1][j], pre[j] - pre[i] - f[i][j-1])
        }
    }

    return f[0][n-1]
}