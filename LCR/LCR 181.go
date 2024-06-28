func reverseMessage(message string) string {
    s := strings.Split(message, " ")
    builder := strings.Builder{}
    n := len(s)
    for i := n - 1; i >= 0; i-- {

        if len(s[i]) > 0 && s[i] != " " {
            builder.WriteString(s[i])
            builder.WriteString(" ")
        }
    }

    ans := builder.String()
    if len(ans) > 0 && ans[len(ans)-1] == ' ' {
        return ans[:len(ans)-1]
    }

    return ans
}
