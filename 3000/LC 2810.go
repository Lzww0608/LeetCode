func finalString(s string) string {
    str := []byte(s)

    reverse := func(i, j int) {
        for i < j {
            str[i], str[j] = str[j], str[i]
            i++
            j--
        }
    }

    for i := range str {
        if str[i] == 'i' {
            reverse(0, i-1)
        }
    }

    builder := strings.Builder{}
    for i := range str {
        if str[i] != 'i' {
            builder.WriteByte(str[i])
        }
    }

    return builder.String()
}
