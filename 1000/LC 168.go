func convertToTitle(columnNumber int) string {
    builder := strings.Builder{}
    for columnNumber > 0 {
        columnNumber--
        t := byte((columnNumber) % 26 + 'A')
        columnNumber = columnNumber / 26
        builder.WriteByte(t)
    }
    ans := builder.String()
    builder.Reset()
    for i := len(ans) - 1; i >= 0; i-- {
        builder.WriteByte(ans[i])
    }
    return builder.String()
}
