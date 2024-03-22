func isHappy(n int) bool {
    m := map[int]bool{}
    for {
        m[n] = true
        x := 0
        for n > 0 {
            t := n % 10
            n /= 10
            x += t * t
        }
        if x == 1 {
            return true
        }
        if m[x] {
            return false
        }
        n = x
    }

    return false
}