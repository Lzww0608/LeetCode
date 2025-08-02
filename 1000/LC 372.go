const MOD int = 1337
func superPow(a int, b []int) int {
    ans := 1

    for i := len(b) - 1; i >= 0; i-- {
        x := b[i]
        if x != 0 {
            ans = ans * (quickPow(a, x)) % MOD
        }
        a = quickPow(a, 10)
    }

    return ans
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