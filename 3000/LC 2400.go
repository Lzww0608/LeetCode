const MOD int = 1_000_000_007

// 用于计算 x 的阶乘的逆元，使用快速幂计算 x^(MOD-2) % MOD
func modInverse(x int) int {
    res := 1
    exp := MOD - 2
    base := x
    for exp > 0 {
        if exp & 1 == 1 {
            res = res * base % MOD
        }
        base = base * base % MOD
        exp >>= 1
    }
    return res
}

// 计算组合数 C(n, k) % MOD
func combination(n int, k int) int {
    if k > n {
        return 0
    }
    numerator := 1
    for i := 0; i < k; i++ {
        numerator = numerator * (n - i) % MOD
    }

    denominator := 1
    for i := 1; i <= k; i++ {
        denominator = denominator * i % MOD
    }

    return numerator * modInverse(denominator) % MOD
}

func numberOfWays(a int, b int, k int) int {
    if k < abs(a - b) {
        return 0
    }

    d := abs(b - a)
    dif := k - d 
    if dif % 2 == 1 {
        return 0
    }
    dif >>= 1

    return combination(k, dif)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
