func unSuitability(operate []int) int {
    mx := slices.Max(operate)
    mx = mx * 2 + 1
    
    n := len(operate)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, mx)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }

    f[0][operate[0]] = operate[0]
    for i := 1; i < n; i++ {
        x := operate[i]
        for j := 0; j < mx; j++ {
            if j + x < mx {
                f[i][j + x] = min(f[i][j + x], max(f[i - 1][j], j + x))
            }

            if j - x >= 0 {
                f[i][j - x] = min(f[i][j - x], f[i - 1][j])
            } else {
                f[i][0] = min(f[i][0], f[i - 1][j] + x - j)
            }
        }
    }

    return slices.Min(f[n - 1])
}