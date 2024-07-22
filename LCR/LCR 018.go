func isPalindrome(s string) bool {
    a := []rune{}
    for _, c := range s {
        if unicode.IsLower(c) || unicode.IsDigit(c) {
            a = append(a, c)
        } else if unicode.IsUpper(c) {
            a = append(a, rune(c + 32))
        }
    }

    if len(a) == 0 {
        return true
    }

    for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
        if a[i] != a[j] {
            return false
        }
    }

    return true
}