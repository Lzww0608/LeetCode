func minWindow(s string, t string) string {
    m := map[int]int{}
    ms := map[int]int{}
    for _, c := range t {
        m[int(c-'A')]++
    }
    l, r, i, j := 0, 0, -1, -1
    n := len(s)
    check := func() bool {
        for k, v := range m {
            if ms[k] < v {
                return false
            }
        }
        return true
    }
    for r < n {
        ms[int(s[r]-'A')]++
        for r - l + 1 >= len(t) && check() {
            if i == -1 || (r - l) < (j - i) {
                i, j = l, r
            } 
            ms[int(s[l]-'A')]--
            l++
        }
        r++
    }
    if i == -1 {
        return ""
    }
    return s[i:j+1]
}