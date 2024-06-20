func getMaximumConsecutive(coins []int) int {
    ans := 1
    sort.Ints(coins)
    for i := 0; i < len(coins) && coins[i] <= ans; i++ {
        ans += coins[i]
    }
    return ans
}
