
import "math"
func maxCoins(a []int, b []int) int64 {
    n := len(a)
    f := [2]int{}
    g := [2]int{}
    for i := range f {
        f[i], g[i] = math.MinInt32, math.MinInt32
    }

    ans := math.MinInt
    for i := 0; i < n; i++ {
        x, y := a[i], b[i]
        g[1], f[1] = max(g[1], f[0], 0) + y, max(f[1], g[1]) + x
        f[0] = max(x, f[0] + x)
        ans = max(ans, f[0], f[1], g[1])
    }

    return int64(ans)
}