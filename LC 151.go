func reverseWords(s string) string {
    a := []string{}
    builder := strings.Builder{}
    for i := range s {
        if s[i] == ' '{
            if builder.Len() > 0 {
                a = append(a, builder.String())
                builder.Reset()
            }
        } else {
            builder.WriteByte(s[i])
        }
    }
    if builder.Len() > 0 {
        a = append(a, builder.String())
        builder.Reset()
    }
    for i := len(a) - 1; i >= 0; i-- {
        builder.WriteString(a[i])
        if i > 0 {
            builder.WriteString(" ")
        }
    }

    return builder.String()
}