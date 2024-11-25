func maxPoints(points [][]int) int64 {
    n := len(points[0])
    f := make([]int, n)
    pre := make([]int, n + 1)
    suf := make([]int, n + 1)
    pre[0], suf[n] = math.MinInt32, math.MinInt32
    for i := range points {
        for j, x := range points[i] {
            if i == 0 {
                f[j] = x
            } else {
                f[j] = x - j + pre[j+1]
                f[j] = max(f[j], x + j + suf[j])
            }
        }
        //fmt.Println(pre, suf, f)
        for j := 1; j <= n; j++ {
            pre[j] = max(pre[j-1], f[j-1] + j - 1)
        }
        for j := n - 1; j >= 0; j-- {
            suf[j] = max(suf[j+1], f[j] - j)
        }
    }

    return int64(slices.Max(f))
}