func minPrice(prices []int, discounts []int) (ans float64) {
    sort.Ints(prices)
    sort.Ints(discounts)
    m, n := len(prices), len(discounts)
    for i, j := m - 1, n - 1; i >= 0; i-- {
        cur := float64(prices[i])
        if j >= 0 {
            cur = (cur * float64(100 - discounts[j])) / 100.0 
            j--
        }
        ans += cur
    }

    return
}