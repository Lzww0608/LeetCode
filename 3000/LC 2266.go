const MOD int = 1_000_000_007
func countTexts(s string) (ans int) {
    n := len(s)

    f := make([]int, max(n + 1, 4))
    g := make([]int, max(n + 1, 4))
    f[0], f[1], f[2], f[3] = 1, 1, 2, 4
    g[0], g[1], g[2], g[3] = 1, 1, 2, 4
    
    for i := 4; i <= n; i++ {
        f[i] = (f[i-1] + f[i-2] + f[i-3]) % MOD
        g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % MOD
    }

    cnt := 0
    ans = 1
    for i := range s {
        cnt++
        if i == n - 1 || s[i+1] != s[i] {
            if s[i] != '7' && s[i] != '9' {
                ans *= f[cnt] % MOD
            } else {
                ans *= g[cnt] % MOD
            }
            ans %= MOD
            cnt = 0
        } 
    }

    return
}