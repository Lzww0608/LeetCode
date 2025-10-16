func maximumProfit(prices []int, k int) int64 {
    f := make([][3]int, k + 1)
    for j := range k + 1 {
        f[j] = [3]int{0, math.MinInt32, 0}
    }

    for _, price := range prices {
        for j := k; j > 0; j-- {
            f[j][0] = max(f[j][0], f[j][1] + price, f[j][2] - price)
            f[j][1] = max(f[j][1], f[j - 1][0] - price)
            f[j][2] = max(f[j][2], f[j - 1][0] + price)
        }
    }

    return int64(f[k][0])
}