const MOD int = 1_000_000_007
const N int = 100_001

var F, invF [N]int

func init() {
    F[0] = 1
    for i := 1; i < N; i++ {
        F[i] = F[i-1] * i % MOD
    }
    invF[N-1] = quickPow(F[N-1], MOD - 2)
    for i := N - 1; i > 0; i-- {
        invF[i-1] = invF[i] * i % MOD
    }
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

func countGoodArrays(n int, m int, k int) int {
    // C(n - 1, k) * m * (m - 1) ^ (n - k - 1)

    return F[n - 1] * invF[k] % MOD * invF[n - 1 - k] % MOD * m % MOD * quickPow(m - 1, n - k - 1) % MOD
}