func coinChange(coins []int, amount int) int {
    f := make([]int, amount + 1)
    for i := range f {
        f[i] = 0x3f3f3f3f
    }
    f[0] = 0

    for i := len(coins) - 1; i >= 0; i-- {
        for j := coins[i]; j <= amount; j++ {
            f[j] = min(f[j-coins[i]] + 1, f[j])
        }
    }

    if f[amount] == 0x3f3f3f3f {
        return -1
    }
    return f[amount]
}
