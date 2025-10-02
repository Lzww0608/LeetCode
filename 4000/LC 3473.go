const INF = 1 << 30 

func maxSum(nums []int, k int, m int) int {
    n := len(nums)
    f := make([][]int, n + 1)
    pre := make([]int, n + 1)
    for i := range n {
        pre[i + 1] = pre[i] + nums[i]
    }
    for i := range f {
        f[i] = make([]int, k + 1)
        for j := 1; j <= k; j++ {
            f[i][j] = -INF
        }
    }

    for j := 1; j <= k; j++ {
        mx := -INF
        for i := j * m; i <= n; i++ {
            p := i - m
            if f[p][j - 1] != -INF {
                mx = max(mx, f[p][j - 1] - pre[p])
            }

            f[i][j] = f[i - 1][j]
            if mx != -INF {
                f[i][j] = max(f[i][j], pre[i] + mx)
            }
        }
    }

    return f[n][k]
}