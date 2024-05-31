func removeDuplicates(s string, k int) string {
    type pair struct {
        c byte 
        n int
    }
    st := []pair{}
    for i := range s {
        if len(st) > 0 && st[len(st)-1].c == s[i] && st[len(st)-1].n == k - 1 {
            for len(st) > 0 && st[len(st)-1].c == s[i] {
                st = st[:len(st)-1]
            }
        } else {
            if len(st) > 0 && st[len(st)-1].c == s[i] {
                st = append(st, pair{s[i], st[len(st)-1].n + 1})
            } else {
                st = append(st, pair{s[i], 1})
            }
        }
    }

    builder := strings.Builder{}
    for _, p := range st {
        builder.WriteByte(p.c)
    }
    return builder.String()
}