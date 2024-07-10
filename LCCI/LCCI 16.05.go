func trailingZeroes(n int) (ans int) {
    for n >= 5 {
        n /= 5
        ans += n
    }

    return
}