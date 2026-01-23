func decimalRepresentation(n int) (ans []int) {
    for i := 1; n > 0; i *= 10 {
        if n % 10 != 0 {
            ans = append(ans, n % 10 * i)
        }
        n /= 10
    }
    slices.Reverse(ans)
    return
}