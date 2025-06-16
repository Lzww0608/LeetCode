
import "math"

func minimumChanges(s string, k int) int {
    n := len(s)
    v := make([][]int, n)
    for i := 0; i < n; i++ {
        v[i] = make([]int, n)
        for j := range v[i] {
            v[i][j] = math.MaxInt32
        }
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            d := j - i + 1
            for t := 1; t < d; t++ {
                if d % t != 0 {
                    continue
                }
                cnt := 0
                for p := 0; p < d; p++ {
                    q := p % t + (d / t - p / t - 1) * t 
                    if p >= q {
                        break
                    }
                    if s[i+q] != s[i+p] {
                        cnt++
                    }
                }
                v[i][j] = min(v[i][j], cnt)
            }
        }
    }

    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, k + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
        f[i][1] = v[0][i]
    }

    for i := 1; i < n; i++ {
        for j := 2; j <= min(k, i + 1); j++ {
            for t := j - 1; t < i; t++ {
                f[i][j] = min(f[i][j], f[t][j-1] + v[t+1][i])
            }
        }
    }

    return f[n-1][k]
}