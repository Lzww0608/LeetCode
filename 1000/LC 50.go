func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }

    f := false
    if n < 0 {
        n = -n
        f = true
    }

    var res float64 = 1.0
    for n > 0 {
        if n & 1 == 1 {
            res *= x
        }
        n >>= 1
        x *= x
    }

    if f {
        return 1.0 / res
    }
    
    return res
}
