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

func countDistinctStrings(s string, k int) int {
    n := len(s)
    return quickPow(2, n - k + 1)
}