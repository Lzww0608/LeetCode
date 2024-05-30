func maxSales(sales []int) int {
    pre, ans := 0, math.MinInt32
    for _, x := range sales {
        pre = max(pre + x, x)
        ans = max(ans, pre)
    }

    return ans
}