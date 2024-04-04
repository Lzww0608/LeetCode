func reverseOnlyLetters(s string) string {
    check := func(c byte) bool {
        if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
            return true
        }
        return false
    }

    ans := []byte(s)
    n := len(ans)
    l, r := 0, n - 1
    for l < r {
        for l < r && !check(ans[l]) {
            l++
        }
        for l < r && !check(ans[r]) {
            r--
        }
        ans[l], ans[r] = ans[r], ans[l]
        l, r = l + 1, r - 1
    }

    return string(ans)
}