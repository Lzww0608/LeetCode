func parseTernary(s string) string {
    n := len(s)
    st := []byte{}
    for i := n - 1; i >= 0; i-- {
        if len(st) > 0 && st[len(st)-1] == '?' {
            // pop '?'
            st = st[:len(st)-1]

            trueResult := st[len(st)-1]
            falseResult := st[len(st)-2]
            st = st[:len(st)-2]

            if s[i] == 'T' {
                st = append(st, trueResult)
            } else {
                st = append(st, falseResult)
            }
        } else if s[i] != ':' {
            st = append(st, s[i])
        }
    } 

    return string(st[0])
}