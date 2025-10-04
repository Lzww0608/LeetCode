func longestPalindrome(s string, t string) int {
    n, m := len(s), len(t)

    ans := 1
    for i := range n {
        for j := i; j <= n; j++ {
            a := s[i:j]
            for p := range m {
                for q := p; q <= m; q++ {
                    b := t[p:q]
                    if ans < j - i + q - p && check(a + b) {
                        ans = j - i + q - p
                    }
                }
            }
        }
    }

    return ans
}

func check(s string) bool {
    n := len(s)
    for i := 0; i * 2 < n; i++ {
        if s[i] != s[n - i - 1] {
            return false
        }
    }

    return true
}