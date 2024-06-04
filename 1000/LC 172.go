func trailingZeroes(n int) (ans int) {
    for i := 5; i <= n; i *= 5 {
        ans += n / i
    }
    return 
}