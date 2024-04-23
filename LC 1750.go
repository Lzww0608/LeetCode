func minimumLength(s string) int {
    n := len(s)
    l, r := 0, n - 1
    for l < r {
        if s[l] != s[r] {
            break
        }
        a := s[l]
        for l <= r && s[l] == a {
            l++
        }
        for l <= r && s[r] == a {
            r--
        }
    }

    return r - l + 1
}