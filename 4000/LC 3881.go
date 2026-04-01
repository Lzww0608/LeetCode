const MOD int = 1_000_000_007
func countVisiblePeople(n int, pos int, k int) int {
    a, b := 1, 1
    for i := range k {
        a = a * (n - 1 - i) % MOD
        b = b * (i + 1) % MOD
    }

    return 2 * a * quickPow(b, MOD - 2) % MOD
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        r >>= 1
        a = a * a % MOD
    }

    return res
}