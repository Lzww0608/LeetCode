func fib(n int) int {
    const MOD int = 1e9 + 7
    if n < 2 {
        return n
    } 
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, (a + b) % MOD
    }
    return b % MOD
}