
import "math"
func connectTwoGroups(cost [][]int) int {
    m, n := len(cost), len(cost[0])
    min_cost := make([]int, n)
    for i := range min_cost {
        min_cost[i] = math.MaxInt32
    }
    for j := 0; j < n; j++ {
        for _, v := range cost {
            min_cost[j] = min(min_cost[j], v[j])
        }
    }

    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, 1 << n)
    }

    for j := 0; j < (1 << n); j++ {
        for k := 0; k < n; k++ {
            if (j >> k) & 1 == 1 {
                f[0][j] += min_cost[k];
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < (1 << n); j++ {
            cur := math.MaxInt32
            for k := 0; k < n; k++ {
                cur = min(cur, f[i][j & ^(1 << k)] + cost[i][k])
            }
            f[i+1][j] = cur
        }
    }

    return f[m][(1<<n)-1]
}