func replaceSpaces(S string, length int) string {
    builder := strings.Builder{}
    for i := 0; i < length; i++ {
        if S[i] == ' ' {
            builder.WriteString("%20")
        } else {
            builder.WriteByte(S[i])
        }
    }

    return builder.String()
}