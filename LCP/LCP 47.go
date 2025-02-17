const MOD int = 1_000_000_007
func securityCheck(capacities []int, k int) int {
    f := make([]int, k + 1)
    f[0] = 1
    for _, x := range capacities {
        x -= 1
        for i := k; i >= x; i-- {
            f[i] = (f[i] + f[i-x]) % MOD
        }
    }

    return f[k]
}