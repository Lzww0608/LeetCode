func patternMatching(pattern string, value string) bool {
    n, m := len(pattern), len(value)
    if n == 0 && m == 0 {
        return true
    } else if n == 0 {
        return false
    }

    a, b := 0, 0
    for _, c := range pattern {
        if c == 'a' {
            a++
        } else {
            b++
        }
    }

    dA, dB := 0, 0
    if a > 0 {
        dA = m / a
    } 
    if b > 0 {
        dB = m / b
    }

    for i := 0; i <= dA; i++ {
        for j := 0; j <= dB; j++ {
            if i * a + j * b != m {
                continue
            }
            sA, sB := "#", "#"
            idx := 0
            for _, c := range pattern {
                if c == 'a' {
                    if sA == "#" {
                        sA = value[idx:i+idx]
                    } else if sA != value[idx:idx+i] {
                        break
                    }
                    idx += i
                } else {
                    if sB == "#" {
                        sB = value[idx:j+idx]
                    } else if sB != value[idx:idx+j] {
                        break
                    }
                    idx +=j
                }
            }

            if idx == m && sA != sB {
                return true
            }
        }
    }

    return false
}