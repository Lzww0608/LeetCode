func isValid(s string) bool {
    if len(s) % 3 != 0 {
        return false
    }
    st := []byte{}
    for i := range s {
        if len(st) == 0 || s[i] == 'a' || s[i] == 'b' {
            st = append(st, s[i])
        } else {
            if len(st) < 2 || st[len(st) - 1] != 'b' || st[len(st) - 2] != 'a' {
                return false
            }
            st = st[:len(st) - 2]
        }
    }

    return len(st) == 0
}