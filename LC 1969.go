var MOD int = int(1e9 + 7)
func minNonZeroProduct(p int) int {
    return (1 << p - 1) % MOD * pow((1 << p - 2), 1 << (p - 1) - 1) % MOD
}

func pow(a, r int) int {
    res := 1

    for a = a % MOD; r > 0; r, a = r >> 1, a * a % MOD {
        if r & 1 == 1 {
            res = res * a % MOD
        }
    }

    return res % MOD
}