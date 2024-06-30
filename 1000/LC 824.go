func toGoatLatin(sentence string) string {
    str := strings.Split(sentence, " ")
    builder := strings.Builder{}
    for i, s := range str {
        if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' || s[0] == 'A' || s[0] == 'E' || s[0] == 'I' || s[0] == 'O' || s[0] == 'U'  {
            for j := range s {
                builder.WriteByte(s[j])
            }
            builder.WriteString("ma")
        } else {
            for j := 1; j < len(s); j++ {
                builder.WriteByte(s[j])
            }
            builder.WriteByte(s[0])
            builder.WriteString("ma")
        }

        for j := 0; j <= i; j++ {
            builder.WriteByte('a')
        }

        if i < len(str) - 1 {
            builder.WriteByte(' ')
        }
    }

    return builder.String()
}
