func countAndSay(n int) string {
    builder := strings.Builder{}
    s := "1"
    for i := 1; i < n; i++ {
        builder.Reset()
        cnt := 1
        for j := 1; j < len(s); j++ {
            if s[j] == s[j-1] {
                cnt++
            } else {
                builder.WriteString(strconv.Itoa(cnt))
                builder.WriteByte(s[j-1])
                cnt = 1
            }
        }
        builder.WriteString(strconv.Itoa(cnt))
        builder.WriteByte(s[len(s)-1])
        s = builder.String()
    }

    return s
}
