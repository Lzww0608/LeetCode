func maximumLength(s string) (ans int) {
    m := map[byte]map[int]int{}
    l, r, n := 0, 0, len(s)
    for r < n {
        for r < n && s[r] == s[l] {
            r++
        }
        if _, ok := m[s[l]]; !ok {
            m[s[l]] = make(map[int]int)
        }
        m[s[l]][r - l]++
        l = r
    }
    for _, v := range m {
        mx := 1
        for x := range v {
            mx = max(x, mx)
        }
        if v[mx] >= 3 {
            ans = max(ans, mx)
        } else if v[mx-1] > 0 || v[mx] >= 2 {
            ans = max(ans, mx - 1)
        } else {
            ans = max(ans, mx - 2)
        }
    }
    if ans == 0 {
        return -1
    }
    return 
}