const MOD int = 1_000_000_007

func makeStringSorted(s string) (ans int) {
    n := len(s)

    F := make([]int, n + 1)
    invF := make([]int, n + 1)
    F[0], invF[0] = 1, 1
    for i := 1; i < n; i++ {
        F[i] = F[i-1] * i % MOD
        invF[i] = quickPow(F[i], MOD - 2)
    }

    cnt := [26]int{}
    for _, c := range s {
        cnt[c - 'a']++
    }

    for i := 0; i < n - 1; i++ {
        rk := 0 
        for j := 0; j < int(s[i] - 'a'); j++ {
            rk += cnt[j]
        }
        cur := rk * F[n-i-1] % MOD
        for j := 0; j < 26; j++ {
            cur = cur * invF[cnt[j]] % MOD
        }
        ans = (ans + cur) % MOD
        cnt[int(s[i] - 'a')]--
    }

    return 
}


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