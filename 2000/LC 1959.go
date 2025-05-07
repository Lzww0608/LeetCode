
import "math"
func minSpaceWastedKResizing(nums []int, k int) int {
    n := len(nums)

    g := make([][]int, n)
    for i := 0; i < n; i++ {
        g[i] = make([]int, n)
        mx := math.MinInt32 
        sum := 0
        for j := i; j < n; j++ {
            mx = max(mx, nums[j])
            sum += nums[j]
            g[i][j] = mx * (j - i + 1) - sum 
        }
    }

    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, k + 2)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
        for j := 1; j <= k + 1; j++ {
            for p := 0; p <= i; p++ {
                if p == 0 {
                    f[i][j] = min(f[i][j], g[p][i])
                } else {
                    f[i][j] = min(f[i][j], g[p][i] + f[p-1][j-1])
                }
            }
        }
    }

    return f[n-1][k+1]
}