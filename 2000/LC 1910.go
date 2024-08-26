func removeOccurrences(s string, part string) string {
    st := []byte{}
    n := len(part)
    for i := range s {
        if s[i] == part[n-1] {
            if len(st) >= n-1 && string(st[len(st)-n+1:]) == part[:n-1] {
                st = st[:len(st)-n+1]
            } else {
                st = append(st, s[i])
            }
        } else {
            st = append(st, s[i])
        }
    }

    return string(st)
}