const N int = 26
const MOD int = 1_000_000_007
const M int = 10_001
var F [M]int 

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

func init() {
    F[0] = 1
    for i := 1; i < M; i++ {
        F[i] = F[i-1] * i % MOD
    }
}

func comb(a, b int) int {
    return F[a] * quickPow(F[b] * F[a - b] % MOD, MOD - 2) % MOD 
}

func countGoodSubsequences(s string) (ans int) {
    cnt := [N]int{}
    mx := 0
    for _, c := range s {
        cnt[c - 'a']++
        mx = max(mx, cnt[c - 'a'])
    }

    for i := 1; i <= mx; i++ {
        cur := 1
        for  _, x := range cnt {
            if x >= i {
                cur = cur * (comb(x, i) + 1) % MOD 
            }
        }

        ans = (ans + cur - 1) % MOD
    }

    return 
}