func toLowerCase(s string) string {
    builder := strings.Builder{}

    for i := range s {
        if s[i] >= 'A' && s[i] <= 'Z' {
            builder.WriteByte(byte(s[i] + 32))
        } else {
            builder.WriteByte(s[i])
        }

    }

    return builder.String()

}