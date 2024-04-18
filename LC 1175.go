const MOD int = 1e9 + 7
func numPrimeArrangements(n int) int {
    cnt := 0
    f := make([]int, n + 1)
    for i := 2; i <= n; i++ {
        if f[i] == 1 {
            continue
        }
        f[i] = 1
        cnt++
        for j := i * i; j <= n; j += i {
            f[j] = 1
        }
    }

    factorial := func(x int) int {
        res := 1
        for x > 1 {
            res = res * x % MOD
            x--
        }
        return res
    }

    ans := factorial(cnt) * factorial(n - cnt) % MOD

    return ans
}