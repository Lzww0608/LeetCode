func largestEven(s string) string {
    i := len(s) - 1
    for i >= 0 && s[i] == '1' {
        i--
    }

    if i < 0 {
        return ""
    }

    return s[:i + 1]
}