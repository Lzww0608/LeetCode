func partitionString(s string) []string {
    m := make(map[string]bool)
    n := len(s)
    ans := make([]string, 0, n)

    var sb strings.Builder
    for i := range s {
        sb.WriteByte(s[i])
        if !m[sb.String()] {
            m[sb.String()] = true
            ans = append(ans, sb.String())
            sb.Reset()
        }
    }

    return ans
}