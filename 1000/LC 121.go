func maxProfit(prices []int) int {
    ans, minPrice := 0, math.MaxInt32
    for _, x := range prices {
        minPrice = min(minPrice, x)
        ans = max(ans, x - minPrice)
    }
    return ans
}
