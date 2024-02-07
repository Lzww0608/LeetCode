func new21Game(n int, k int, maxPts int) float64 {
    if k == 0 || k + maxPts <= n {
        return 1.0
    }
    dp := make([]float64, n + 1)
    dp[0] = 1.0
    var ans float64 = 0.0
    var windowSum float64 = 1.0
    for i := 1; i <= n; i++ {
        dp[i] = windowSum / float64(maxPts)
        if i >= k {
            ans += dp[i]
        } else {
            windowSum += dp[i]
        }
        if i >= maxPts {
            windowSum -= dp[i-maxPts]
        }
    }
    return ans
}