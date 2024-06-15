func interpret(command string) string {
    i, n := 0, len(command)

    builder := strings.Builder{}
    for i < n {
        if command[i] == 'G' {
            builder.WriteByte('G')
        } else if command[i] == '(' && command[i+1] == ')' {
            builder.WriteByte('o')
            i++
        } else {
            for i < n && command[i] != ')' {
                i++
            }
            builder.WriteString("al")
        }
        i++
    }

    return builder.String()
}
