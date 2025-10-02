func processStr(s string, k int64) byte {
    d := int64(0)
    for  _, c := range s {
        if c == '*' {
            d = max(d - 1, 0)
        } else if c == '#' {
            d *= 2
        } else if c != '%' {
            d++
        }
    }

    if k >= d {
        return '.'
    }

    for i := len(s) - 1; i >= 0; i-- {
        c := s[i]
        if c == '*' {
            d++
        } else if c == '#' {
            d /= 2 
            if k >= d {
                k -= d
            }
        } else if c == '%' {
            k = d - 1 - k
        } else {
            d--
            if k == d {
                return c
            }
        }
    }

    return '.'
}