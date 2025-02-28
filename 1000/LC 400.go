func findNthDigit(n int) int {
    ans := 0
    p, tmp := 0, 0
    for ans < n {
        tmp = ans 
        ans += 9 * quickPow(10, p) * (p + 1)
        p++
    }

    if ans == n {
        return 9
    }

    d := n - tmp
    a := d / p
    b := d % p
    res := quickPow(10, p - 1) + a - 1
    if b == 0 { 
        return res % 10
    }

    s := strconv.Itoa(res + 1)
    return int(s[b-1] - '0')
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