func isPowerOfThree(n int) bool {
    for i := 1; i <= n; i *= 3 {
        if i == n {
            return true
        }
    }
    return false
}
