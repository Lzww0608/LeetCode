func breakPalindrome(palindrome string) string {
    if len(palindrome) == 1 {
        return ""
    }
    s := []byte(palindrome)
    n := len(s)
    for i := range s {
        if s[i] > 'a' && (n & 1 == 0 || i != n / 2) {
            s[i] = 'a'
            return string(s)
        }
    }

    s[n-1] += 1
    return string(s)
}