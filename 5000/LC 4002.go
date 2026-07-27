const MOD int = 1_000_000_007

// C(n - 1, k - 1)
// y = 2 * x - 1, sumY = n
// (n + k) / 2 = sumX, (n + k) must be even
// C((n + k) / 2 - 1, k - 1)

const N int = 500_001
var F [N]int 
var invF [N]int
func init() {
    F[1] = 1
    for i := 2; i < N; i++ {
        F[i] = F[i - 1] * i % MOD
    }

    invF[N - 1] = quickPow(F[N - 1], MOD - 2)
    for i := N - 1; i > 0; i-- {
        invF[i - 1] = invF[i] * i % MOD
    }
}

func countValidSequences(n int, k int) int {
    a := F[n - 1] * invF[k - 1] % MOD * invF[(n - 1) - (k - 1)] % MOD

    if (n + k) & 1 == 1 {
        return a
    }
    x, y := (n + k) / 2 - 1, k - 1
    b := F[x] * invF[y] % MOD * invF[x - y] % MOD 
    return (a - b + MOD) % MOD
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