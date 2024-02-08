func backspaceCompare(s string, t string) bool {
    var a, b strings.Builder
    back := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '#' {
            back++
        } else if back > 0 {
            back--
        } else {
            a.WriteByte(s[i])
        }
    }
    back = 0
    for i := len(t) - 1; i >= 0; i-- {
        if t[i] == '#' {
            back++
        } else if back > 0 {
            back--
        } else {
            b.WriteByte(t[i])
        }
    }
    return a.String() == b.String()
}