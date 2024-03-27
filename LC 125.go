func isPalindrome(s string) bool {
    builder := strings.Builder{}
    
    for i := range s {
        if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
            builder.WriteByte(s[i])
        } else if s[i] >= 'A' && s[i] <= 'Z' {
            builder.WriteByte(byte(s[i] + 32))
        }
    }

    str := builder.String()
    for l, r := 0, len(str) - 1; l < r; l, r = l + 1, r - 1 {
        if str[l] != str[r] {
            return false
        }
    }

    return true
}