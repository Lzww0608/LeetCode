func removeOuterParentheses(input string) string {
    s := []byte(input)
    cnt := 0
    builder := strings.Builder{}
    for i := range s {
        if s[i] == ')' {
            cnt--
        } 
        if cnt >= 1 {
            builder.WriteByte(s[i])
        } 
        if s[i] == '(' {
            cnt++
        }
    }

    return builder.String()
}