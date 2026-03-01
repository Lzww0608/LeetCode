func mergeCharacters(s string, k int) string {
    st := []byte{}
    for i := range s {
        n := len(st)
        f := true
        for j := range min(k, n) {
            if st[n - j - 1] == s[i] {
                f = false
                break
            }
        }

        if f {
            st = append(st, s[i])
        }
    }

    return string(st)
}