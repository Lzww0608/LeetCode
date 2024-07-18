const MOD int = 1_000_000_007
func waysToStep(n int) int {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    a, b, c := 1, 1, 2

    for i := 3; i <= n; i++ {
        a, b, c = b, c, (a + b + c) % MOD
    }

    return c
}
