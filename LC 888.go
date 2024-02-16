func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sum_alice, sum_bob := 0, 0
    m := map[int]bool{}
    for _, x := range aliceSizes {
        sum_alice += x
    }
    for _, x := range bobSizes {
        sum_bob += x
        m[x] = true
    }
    candies := (sum_alice - sum_bob) / 2
    for _, x := range aliceSizes {
        if _, ok := m[x - candies]; ok {
            return []int{x, x - candies}
        }
    }
    return []int{}
}