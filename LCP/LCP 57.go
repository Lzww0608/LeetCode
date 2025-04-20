
import (
	"math"
)
func getMaximumNumber(moles [][]int) int {
    a := [][]int{{0, 1, 1}}
    add := 0
    for _, v := range moles {
        if v[0] == 0 {
            if v[1] == 1 && v[2] == 1 {
                add = 1
            }
        } else {
            a = append(a, v)
        }
    }

    n := len(a)
    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0]
    })
    f := make([]int, n)
    g := make([]int, n)

    for i := 1; i < n; i++ {
        f[i] = math.MinInt32
        for j := i - 1; j >= 0; j-- {
            x, y := a[i][0] - a[j][0], abs(a[i][1] - a[j][1]) + abs(a[i][2] - a[j][2])
            if x >= 4 {
                f[i] = max(f[i], g[j] + 1)
                break
            } else if y <= x {
                f[i] = max(f[i], f[j] + 1)
            }
        }
        g[i] = max(g[i-1], f[i])
    }

    return g[n-1] + add
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x 
}