func getSmallestString(n int, k int) string {
    builder := strings.Builder{}

    for i := 0; i < n; i++ {
        t := max(1, k - (n - i - 1) * 26) - 1
        k -= t + 1
        builder.WriteByte(byte('a' + t))
    }

    return builder.String()
}
