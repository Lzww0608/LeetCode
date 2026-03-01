const MOD int = 1_000_000_007
func sumOfNumbers(l int, r int, k int) (ans int) {
    cnt := quickPow(r - l + 1, k - 1)
    sum := 0
    for i := l; i <= r; i++ {
        sum += i 
    }
    pow10 := (quickPow(10, k) - 1 +  MOD) % MOD
    inv := quickPow(9, MOD - 2)

    ans = sum * cnt % MOD * pow10 % MOD * inv % MOD
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