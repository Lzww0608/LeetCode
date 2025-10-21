func minXor(nums []int, k int) int {
    n := len(nums)

    f := make([][]int, k + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }
    f[0][0] = 0

    for i := 1; i <= k; i++ {
        for j := i; j <= n - (k - i); j++ {
            s := 0
            for l := j - 1; l >= i - 1; l-- {
                s ^= nums[l]
                f[i][j] = min(f[i][j], max(f[i - 1][l], s))
            }
        }
    }

    return f[k][n]
}