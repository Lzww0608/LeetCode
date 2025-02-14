const MOD int = 1_000_000_007
func numberOfWays(n int, x int) int {
    f := make([]int, n + 1)
    a := []int{}
    for i := 1; ;i++ {
        x := quickPow(i, x)
        if x > n {
            break
        }
        a = append(a, x)
    }

    f[0] = 1
    for _, y := range a {
        for i := n; i >= y; i-- {
            f[i] = (f[i] + f[i-y]) % MOD
        }
    }

    return f[n]
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res *= a
        }
        a = a * a
        r >>= 1
    }

    return res
}