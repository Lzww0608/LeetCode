func hasAlternatingBits(n int) bool {
    for n > 0 {
        x := n & 1
        n >>= 1
        if x == n & 1 {
            return false
        }
    }

    return true
}
