func maxSumAfterPartitioning(arr []int, k int) int {
    n := len(arr)
    dp := make([]int, n + 1)
    for i, x := range arr {
        dp[i+1] = dp[i] + x
        mx := x
        for j := 1; j < k && i - j >= 0; j++ {
            mx = max(mx, arr[i-j])
            dp[i+1] = max(dp[i+1], dp[i-j] + (j + 1) * mx)
        }
        //fmt.Println(i, dp[i+1])
    }
    return dp[n]
}
