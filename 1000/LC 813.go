func largestSumOfAverages(nums []int, k int) (ans float64) {
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i+1] = pre[i] + x
    }

    f := make([][]float64, n + 1)
    for i := range f {
        f[i] = make([]float64, k + 1)
    }
    

    for i := 1; i <= n; i++ {
        f[i][1] = float64(pre[i]) / float64(i)
        for j := 2; j <= k && j <= i; j++ {
            for l := 1; l < i; l++ {
                f[i][j] = max(f[i][j], f[l][j-1] + float64(pre[i] - pre[l]) / float64(i - l))
            }
        }
    }

    return f[n][k]
}
