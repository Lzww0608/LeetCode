func hasSpecialSubstring(s string, k int) bool {
    n := len(s)

    cnt := 0
    for i := range n {
        if i > 0 && s[i] == s[i - 1] {
            cnt++
        } else {
            cnt = 1
        }

        if cnt == k {
            if (i - cnt < 0 || s[i - cnt] != s[i]) && (i + 1 >= n || s[i + 1] != s[i]) {
                return true
            }
        }
    }

    return false
}