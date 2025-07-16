const MOD int = 1_000_000_007

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }

    return res
}

func numberOfSequence(n int, sick []int) (ans int) {
    m := len(sick)
    fac := make([]int, n + 1)
    fac[0] = 1
    for i := 1; i <= n; i++ {
        fac[i] = fac[i-1] * i % MOD
    }
    ans = fac[n-m]


    cnt := 0
    if sick[0] > 0 {
        ans = ans * quickPow(fac[sick[0]], MOD - 2) % MOD
    }
    if n - 1 - sick[m-1] > 0 {
        ans = ans * quickPow(fac[n - 1 - sick[m-1]], MOD - 2) % MOD
    }
    for i := 0; i < m - 1; i++ {
        d := sick[i+1] - sick[i] - 1
        if d <= 1 {
            continue
        }
        cnt += d - 1
        ans = ans * quickPow(fac[d], MOD - 2) % MOD
    }

    return ans * quickPow(2, cnt) % MOD
}