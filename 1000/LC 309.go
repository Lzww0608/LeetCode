func maxProfit(prices []int) int {
    n := len(prices)
    f := make([][3]int, n + 1)
    f[1][0] = -prices[0]

    for i := 1; i < n; i++ {
        x := prices[i]
        f[i+1][0] = max(f[i][0], f[i][2] - x)
        f[i+1][1] = max(f[i][1], f[i][0] + x)
        f[i+1][2] = max(f[i][2], f[i][1])
    }
    
    return max(f[n][1], f[n][2])
}