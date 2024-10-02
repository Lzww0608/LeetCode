import (
    "math"
    "sort"
)
func minDistance(houses []int, k int) int {
    sort.Ints(houses)
    n := len(houses)
    dis := make([][]int, n)

    for i := 0; i < n; i++ {
        dis[i] = make([]int, n)
        for j := i; j < n; j++ {
            mid := i + ((j - i) >> 1)
            for k := i; k <= j; k++ {
                dis[i][j] += int(math.Abs(float64(houses[mid] - houses[k])))
            }
        }
    }

    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, k + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }
    f[0][0] = 0

    for i := 1; i <= n; i++ {
        for j := 1; j <= k; j++ {
            for m := 0; m < i; m++ {
                f[i][j] = min(f[i][j], f[m][j-1] + dis[m][i-1])
            }
        }
    }

    return f[n][k]
}