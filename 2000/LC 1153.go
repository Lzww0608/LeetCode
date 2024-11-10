func canConvert(s string, t string) bool {
    if s == t {
        return true
    }
    cnt := map[int]struct{}{}
    for i := range t {
        x := int(t[i] & 31) - 1
        cnt[x] = struct{}{}
    }
    if len(cnt) == 26 {
        return false
    }

    pos := make([][]int, 26)
    for i := range s {
        x := int(s[i] & 31) - 1
        pos[x] = append(pos[x], i)
    }

    for _, v := range pos {
        for i := 1; i < len(v); i++ {
            if t[v[i]] != t[v[0]] {
                return false
            }
        }
    }

    return true
}