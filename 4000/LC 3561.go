func check(a, b byte) bool {
    if a > b {
        return check(b, a)
    }

    return a + 1 == b || a == 'a' && b == 'z'
}

func resultingString(s string) string {
    st := []byte{}

    for i := range s {
        if len(st) == 0 || !check(st[len(st) - 1], s[i]) {
            st = append(st, s[i])
        } else {
            st = st[:len(st) - 1]
        }
    }

    return string(st)
}