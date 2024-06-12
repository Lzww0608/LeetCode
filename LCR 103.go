const N int = 0x3f3f3f3f
func coinChange(coins []int, amount int) int {
    f := make([]int, amount + 1)
    for i := range f {
        f[i] = N
    }
    f[0] = 0

    for _, x := range coins {
        for k := x; k <= amount; k++ {
            f[k] = min(f[k], f[k-x] + 1)
        }
    }

    if f[amount] == N {
        return -1
    }

    return f[amount]
}