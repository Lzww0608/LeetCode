func maxTastiness(price []int, tastiness []int, maxAmount int, maxCoupons int) int {
    n := len(price)
    dp := make([][]int, maxAmount + 1)
    for i := range dp {
        dp[i] = make([]int, maxCoupons + 1)
    }

    for k := 0; k < n; k++ {
        x := price[k]
        t := tastiness[k]
        for i := maxAmount; i >= x / 2; i-- {
            for j := maxCoupons; j >= 0; j-- {
                if i >= x {
                    dp[i][j] = max(dp[i][j], dp[i-x][j] + t)
                }
                if j > 0 {
                    dp[i][j] = max(dp[i][j], dp[i-x/2][j-1] + t)
                }
            }
        }
    }
    return dp[maxAmount][maxCoupons]
}