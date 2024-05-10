func maxProfit(prices []int) int {
    n := len(prices)
    f := make([][4]int, n)
    f[0][0], f[0][2] = -prices[0], -prices[0]
    for i := 1; i < n; i++ {
        f[i][0] = max(f[i-1][0], -prices[i])
        f[i][1] = max(f[i-1][1], f[i-1][0] + prices[i])
        f[i][2] = max(f[i-1][2], f[i-1][1] - prices[i])
        f[i][3] = max(f[i-1][3], f[i-1][2] + prices[i])
    }

    return f[n-1][3]

}