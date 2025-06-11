const MOD int = 1_000_000_007
func countSubranges(a []int, b []int) (ans int) {
    n := len(a)
    f := make(map[int]int)
    for i := 0; i < n; i++ {
        cur := make(map[int]int)
        cur[a[i]]++
        cur[-b[i]]++

        for p, v := range f {
            if p > (n - i) * 100 || p < -(n - i) * 100 {
                continue
            }
            cur[p + a[i]] = (cur[p + a[i]] + v) % MOD
            cur[p - b[i]] = (cur[p - b[i]] + v) % MOD
        }

        f = cur
        ans = (ans + f[0]) % MOD
    }

    return
}