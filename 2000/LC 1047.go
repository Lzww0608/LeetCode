func removeDuplicates(s string) string {
    st := []byte{}
    for i := range s {
        if len(st) > 0 && st[len(st)-1] == s[i] {
            st = st[:len(st)-1]
        } else {
            st = append(st, s[i])
        }
    }

    return string(st)
}
