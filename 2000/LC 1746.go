func maxSumAfterOperation(nums []int) (ans int) {
    n := len(nums)
    f := make([][2]int, n + 1)
    for i, x := range nums {
        f[i+1][1] = max(f[i][1] + x,  x * x, f[i][0] + x * x)
        f[i+1][0] = max(f[i][0] + x, x)
        ans = max(ans, f[i+1][0], f[i+1][1])
    }

    return
}