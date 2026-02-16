func almostPalindromic(s string) int {
    n := len(s)
    ans := 0
    check := func(l, r int) {
        for l >= 0 && r < n {
            if s[l] != s[r] {
                break
            }
            l--
            r++
        }

        ans = max(ans, min(n, r - l - 1))
    }    

    for i := 0; i < n * 2 - 1 && ans < n; i++ {
        l, r := i / 2, (i + 1) / 2
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
        }
        check(l - 1, r)
        check(l, r + 1) 
    }

    return ans
}
