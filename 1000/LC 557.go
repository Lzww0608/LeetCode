func reverseWords(s string) string {
    str := strings.Split(s, " ")
    builder := strings.Builder{}
    for i := range str {
        b := []byte(str[i])
        for l, r := 0, len(b) - 1; l < r; l, r = l + 1, r - 1 {
            b[l], b[r] = b[r], b[l]
        }
        builder.WriteString(string(b))
        if i < len(str) - 1 {
            builder.WriteByte(' ')
        }
    }
    return builder.String()
}
