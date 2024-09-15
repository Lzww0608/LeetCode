func maxSizeSlices(slices []int) int {

    solve := func(a []int) int {
        n, k := len(a), (len(a) + 1) / 3

        f := make([][]int, n)
        for i := range f {
            f[i] = make([]int, k + 1)
        } 
        f[0][1], f[1][1] = a[0], max(a[0], a[1])

        for i := 2; i < n; i++ {
            for j := 1; j <= k; j++ {
                f[i][j] = max(f[i-1][j], f[i-2][j-1] + a[i])
            }
        }

        return f[n-1][k]
    }
    n := len(slices)
    return max(solve(slices[:n-1]), solve(slices[1:]))
}