func maxCoins(nums []int) int {
    n := len(nums)
    nums = append(nums, 1)
    arr := append([]int{1}, nums...)
    
    dp := make([][]int, n + 2)
    for i := range dp {
        dp[i] = make([]int, n + 2)
    }

    for d := 3; d <= n + 2; d++ {
        for l := 0; l + d - 1 <= n + 1; l++ {
            r := l + d - 1
            for k := l + 1; k < r; k++ {
                dp[l][r] = max(dp[l][r], dp[l][k] + dp[k][r] + arr[k] * arr[l] * arr[r])
            }
        }
    }

    return dp[0][n+1]
}