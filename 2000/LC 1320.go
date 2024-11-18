
import "math"

const N int = 26
func minimumDistance(s string) int {
    n := len(s)
    f := make([][]int, N)
    for i := range f {
        f[i] = make([]int, N)
    }

    ans := math.MaxInt32
    for k := 1; k < n; k++ {
        tmp := make([][]int, N)
        for i := range tmp {
            tmp[i] = make([]int, N)
            for j := range tmp {
                tmp[i][j] = math.MaxInt32
            }
        }
        for i := 0; i < N; i++ {
            for j := 0; j < N; j++ {
                y := int(s[k-1] & 31) - 1
                if i != y && j != y {
                    continue
                }
                x := int(s[k] & 31) - 1
                tmp[i][x] = min(tmp[i][x], f[i][j] + calc(x, j))
                tmp[x][j] = min(tmp[x][j], f[i][j] + calc(x, i))
                if k == n - 1 {
                    ans = min(ans, tmp[i][x], tmp[x][j])
                }
            }
        }
        f = tmp
    }

    return ans
}

func calc(x, y int) int {
    x, y = min(x, y), max(x, y)
    a, b := x % 6, y % 6
    return y / 6 - x / 6 + abs(b - a)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}