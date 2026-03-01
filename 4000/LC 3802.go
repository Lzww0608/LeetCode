const MOD int = 1_000_000_007
func numberOfWays(n int, a []int) (ans int) {
    m := len(a)
    sum := 0
    for i := range a {
        a[i] = min(a[i], n - 1)
        sum += a[i]
    }
    sort.Ints(a)

    for i, j := 0, m - 1; i < j; {
        if a[i] + a[j] < n {
            sum -= a[i]
            i++
        } else {
            sum -= a[j]
            ans = (ans + sum - (n - a[j] - 1) * (j - i)) % MOD
            j--
        }
    }

    return (ans * 2 % MOD + MOD) % MOD
}