func integerBreak(n int) int {
    if n <= 3 {
        return n - 1
    }
    a := n / 3 
    b := n % 3 
    if b == 0 {
        return quickPow(3, a)
    } else if b == 1 {
        return quickPow(3, a - 1) * 4
    } else {
        return quickPow(3, a) * 2
    }
}


func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res *= a
        }
        a *= a
        r >>= 1
    }
    return res
}