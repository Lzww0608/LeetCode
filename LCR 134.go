func myPow(x float64, n int) float64 {
    if x == 0 || x == 1 {
        return x
    }

    f := false
    if n < 0 {
        f = true
        n = -n
    }

    var ans float64 = 1.0
    for ; n > 0; n >>= 1 {
        if n & 1 == 1 {
            ans *= x
        }
        x *= x
    }

    if f {
        return 1.0 / ans
    }

    return ans
}