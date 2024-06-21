const N int = 1010
const inf float64 = 1e9
const eps float64 = 1e-8
func minSkips(dist []int, speed int, hoursBefore int) int {
    n := len(dist)
    f := [N][N]float64{}

    for i := 1; i < n; i++ {
        t := float64(dist[i-1]) / float64(speed)
        for j := 0; j <= i; j++ {
            f[i][j] = inf
            if j < i {
                f[i][j] = math.Ceil(f[i-1][j] + t - eps)
            }
            if j > 0 {
                f[i][j] = min(f[i][j], f[i-1][j-1] + t)
            }
        }
    }
    t := float64(dist[n-1]) / float64(speed)
    for i := 0; i < n; i++ {
        if f[n-1][i] + t < float64(hoursBefore) + eps {
            return i
        }
    }

    return -1
}
