const MOD int = 1_000_000_007
func numberOfWays(n int) int {
    a := []int{1, 2, 6}
    f := make([]int, n + 1)
    f[0] = 1

    for _, x := range a {
        for i := x; i <= n; i++ {
            f[i] = (f[i] + f[i-x]) % MOD
        }
    }

    if n >= 4 {
        f[n] = (f[n] + f[n-4]) % MOD
    } 
    if n >= 8 {
        f[n] = (f[n] + f[n-8]) % MOD
    }

    return f[n]
}