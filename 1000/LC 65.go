func isNumber(s string) bool {
    n := len(s)
    i := 0
    isNum := false
    if s[i] == '-' || s[i] == '+' {
        i++
    }
    for i < n && isDigit(s[i]) {
        i++
        isNum = true
    }

    if i < n && s[i] == '.' {
        i++
    }
    for i < n && isDigit(s[i]) {
        i++
        isNum = true
    }
    if !isNum {
        return false
    }

    if i < n {
        isNum = false
        if s[i] == 'e' || s[i] == 'E' {
            i++
        } else {
            return false
        }
    }

    if i < n && (s[i] == '-' || s[i] == '+') {
        i++
    }

    for i < n && isDigit(s[i]) {
        i++
        isNum = true
    }

    return i == n && isNum
}

func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}