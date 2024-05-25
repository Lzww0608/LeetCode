const MOD int = 1e9 + 7
func trainWays(n int) int {

    a, b := 0, 1
    for i := 1; i <= n; i++ {
        a, b = b, (a + b) % MOD
    }

    return b
}