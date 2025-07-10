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

func stringCount(n int) int {
    // 1个l，1个t，2个e
    // 缺一个条件， 25 ^ n * 3 + n * 25 ^ (n - 1) 
    // 缺两个条件， 24 ^ n + (24 ^ n + n * 24 ^ (n - 1)) * 2
    // 缺三个条件， 23 ^ n + n * 23 ^ (n - 1)

    a := quickPow(25, n) * 3 % MOD + quickPow(25, n - 1) * n % MOD
    b := quickPow(24, n) % MOD + (quickPow(24, n) + quickPow(24, n - 1) * n) % MOD * 2 % MOD 
    c := quickPow(23, n) + quickPow(23, n - 1) * n % MOD

    return (quickPow(26, n) - a % MOD + MOD + b % MOD - c % MOD + MOD) % MOD
}const MOD int = 1_000_000_007

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

func stringCount(n int) int {
    // 1个l，1个t，2个e
    // 缺一个条件， 25 ^ n * 3 + n * 25 ^ (n - 1) 
    // 缺两个条件， 24 ^ n + (24 ^ n + n * 24 ^ (n - 1)) * 2
    // 缺三个条件， 23 ^ n + n * 23 ^ (n - 1)

    a := quickPow(25, n) * 3 % MOD + quickPow(25, n - 1) * n % MOD
    b := quickPow(24, n) % MOD + (quickPow(24, n) + quickPow(24, n - 1) * n) % MOD * 2 % MOD 
    c := quickPow(23, n) + quickPow(23, n - 1) * n % MOD

    return (quickPow(26, n) - a % MOD + MOD + b % MOD - c % MOD + MOD) % MOD
}