func minLength(s string) int {
    st := []byte{}
    for i := range s {
        if len(st) > 0 && (s[i] == 'B' && st[len(st)-1] == 'A' || s[i] == 'D' && st[len(st)-1] == 'C') {
            st = st[:len(st)-1]
        } else {
            st = append(st, s[i])
        }
    }
    return len(st)
}