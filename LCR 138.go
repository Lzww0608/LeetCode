func validNumber(s string) bool {
    i, n := 0, len(s)
    for i < n && s[i] == ' ' {
        i++
    }
    if i == n {
        return false
    }

    checkUint := func(i *int) bool {
        t := *i
        for *i < n && s[*i] >= '0' && s[*i] <= '9' {
            (*i)++
        }
        return *i > t
    }

    checkInt := func(i *int) bool {
        if *i < n && (s[*i] == '+' || s[*i] == '-') {
            (*i)++
        }
        return checkUint(i)
    }
    
    f := checkInt(&i)
    if i < n && s[i] == '.' {
        i++
        f = checkUint(&i) || f
    }

    if i < n && (s[i] == 'e' || s[i] == 'E') {
        i++
        f = checkInt(&i) && f
    }

    for i < n && s[i] == ' ' {
        i++
    }

    return f && i == n
}