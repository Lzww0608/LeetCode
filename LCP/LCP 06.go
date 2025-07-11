func minCount(coins []int) (ans int) {
    for _, x := range coins {
        ans += x / 2 + x & 1
    }

    return 
}