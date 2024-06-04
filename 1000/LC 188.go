func maxProfit(k int, prices []int) int {
    n := len(prices)
    f := make([][2]int, k)
    for i := 0; i < k; i++ {
        f[i][0] = -prices[0]
    }

    for i := 1; i < n; i++ {
        f[0][1] = max(f[0][1], f[0][0] + prices[i])
        f[0][0] = max(f[0][0], -prices[i])
        for j := k - 1; j > 0; j-- {
            f[j][1] = max(f[j][1], f[j][0] + prices[i])
            f[j][0] = max(f[j][0], f[j-1][1] - prices[i])
        }
    }

    return f[k-1][1]
}
