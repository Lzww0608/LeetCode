func freqAlphabets(s string) string {
    n := len(s)
    builder := strings.Builder{}
    for i := 0; i < n; i++ {
        if i < n - 2 && s[i+2] == '#' {
            t := int(s[i] - '0') * 10 + int(s[i+1] - '0') - 10
            builder.WriteByte(byte('j' + t))
            i += 2
        } else {
            builder.WriteByte(byte(int(s[i] - '1') + 'a'))
        }
    }

    return builder.String()
}
