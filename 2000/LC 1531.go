
import "math"
func getLengthOfOptimalCompression(s string, k int) int {
    n := len(s)
    k = n - k 
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, k + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }
    f[n][k] = 0

    for i := n - 1; i >= 0; i-- {
        for cnt := 0; cnt <= k; cnt++ {
            for j, same := i, 0; j < n; j++ {
                if s[j] == s[i] {
                    same++
                }
                if same + cnt > k {
                    break
                }
                f[i][cnt] = min(f[i][cnt], solve(same) + f[j+1][cnt + same])
            }
            f[i][cnt] = min(f[i][cnt], f[i+1][cnt])
        }
    }

    return f[0][0]
}

func solve(x int) int {
    if x <= 1 {
        return x
    } else if x <= 9 {
        return 2
    } else if x <= 99 {
        return 3
    }
    return 4
}