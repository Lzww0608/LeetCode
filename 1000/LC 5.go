func longestPalindrome(s string) (ans string) {
    n := len(s)
    mx := 0

    check := func(i, j int) {
        l, r := i, j
        for l >= 0 && r < n {
            if s[l] == s[r] {
                l--
                r++
            } else {
                break
            }
        }

        if r - l - 1 > mx {
            mx = r - l - 1
            ans = s[l+1:r]
        }
    }

    for i := range s {
        check(i, i)
        check(i, i + 1)
    }

    return 
}

