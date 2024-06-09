const MOD int = 1e9 + 7
func peopleAwareOfSecret(n int, delay int, forget int) int {
    f := make([]int, n + 1)
    f[1] = 1
    for i := 2; i <= n; i++ {
        s := (f[max(i - delay, 0)] - f[max(i - forget, 0)]) % MOD
        f[i] = (f[i-1] + s) % MOD
    }

    return ((f[n] - f[max(n - forget, 0)]) % MOD + MOD) % MOD
}