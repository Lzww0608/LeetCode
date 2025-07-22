func makeAntiPalindrome(S string) string {
    s := []byte(S)
    n := len(s)
    sort.Slice(s, func(i, j int) bool {
        return s[i] < s[j]
    })

    j := n / 2 + 1
    for i := n / 2; i < n && s[i] == s[n - i - 1]; i++ {
        for j < n && s[j] == s[i] {
            j++
        }

        if j >= n {
            return "-1"
        }

        s[i], s[j] = s[j], s[i]
    }

    return string(s)
}