const MOD int = 1_000_000_007
func maxNiceDivisors(n int) int {
    if n <= 3 {
        return n
    }

    a := n / 3 
    b := n % 3 
    if b == 0 {
        return quickPow(3, a)
    } else if b == 1 {
        return quickPow(3, a - 1) * 4 % MOD
    } else {
        return quickPow(3, a) * 2 % MOD
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