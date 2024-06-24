const MOD int = 1_000_000_007
func waysToChange(n int) int {
    f := make([]int, n + 1)
    for i := range f {
        f[i] = 1
    }

    a := []int{5, 10, 25}

    for _, x := range a {
        for i := x; i <= n; i++ {
            f[i] += f[i-x]
            f[i] %= MOD
        }
    }

    return f[n]
}