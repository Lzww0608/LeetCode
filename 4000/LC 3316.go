
import "math"
func maxRemovals(s string, p string, targetIndices []int) int {
    m, n := len(s), len(p)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        for j := range f[i] {
            f[i][j] = math.MinInt32
        }
    }

    cnt := make(map[int]int)
    for _, x := range targetIndices {
        cnt[x]++
    }
    f[0][0] = 0
    for i := 0; i < m; i++ {
        t := cnt[i]
        f[i+1][0] = f[i][0] + t 
        for j := 0; j < min(i + 1, n); j++ {
            if s[i] == p[j] {
                f[i+1][j+1] = max(f[i][j+1] + t, f[i][j])
            } else {
                f[i+1][j+1] = f[i][j+1] + t 
            }
        }
    }

    return f[m][n]
}