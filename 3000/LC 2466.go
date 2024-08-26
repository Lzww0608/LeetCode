const MOD int = 1_000_000_007
func countGoodStrings(low int, high int, zero int, one int) (ans int) {
    f := make([]int, high + 1)
    f[0] = 1
    for i := min(zero, one); i <= high; i++ {
        if i >= zero {
            f[i] += f[i-zero]
        } 
        if i >= one {
            f[i] += f[i-one]
        }

        f[i] %= MOD
        if i >= low {
            ans = (ans + f[i]) % MOD
        }
    }

    return
}