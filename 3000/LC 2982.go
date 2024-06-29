func maximumLength(s string) (ans int) {
    m := [26][]int{}
    l, r, n := 0, 0, len(s)
    for r < n {
        for r < n && s[r] == s[l] {
            r++
        }
        m[int(s[l] - 'a')] = append(m[int(s[l] - 'a')], r - l)
        l = r
    }
    for _, v := range m {
        if len(v) == 0 {
            continue
        }
        sort.Slice(v, func(i, j int) bool {
            return v[i] > v[j]
        })
        v = append(v, 0, 0)
        ans = max(ans, v[0] - 2, v[2], min(v[0] - 1, v[1]))

    }
    if ans == 0 {
        return -1
    }
    return 
}
