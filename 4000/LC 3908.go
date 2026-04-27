func validDigit(n int, x int) bool {
    if n < 10 {
        return false
    }

    cnt := 0
    for n >= 10 {
        if n % 10 == x {
            cnt++
        }
        n /= 10
    }

    return cnt > 0 && n != x
}