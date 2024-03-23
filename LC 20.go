func isValid(s string) bool {
    m := map[byte]byte {
        ')': '(',
        ']': '[',
        '}': '{',
    }

    st := []byte{}
    for i := range s {
        if len(st) > 0 && m[s[i]] == st[len(st)-1] {
            st = st[:len(st)-1]
        } else {
            st = append(st, s[i])
        }
    }

    return len(st) == 0
}