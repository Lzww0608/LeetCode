func strongPasswordChecker(s string) int {
    var misLowChar, misDigit, misUpperChar bool = true, true, true

    for _, c := range s {
        if unicode.IsDigit(c) {
            misDigit = false
        } else if unicode.IsLower(c) {
            misLowChar = false
        } else if unicode.IsUpper(c) {
            misUpperChar = false
        }
    }

    misNum := 0
    if misLowChar {
        misNum++
    }
    if misDigit {
        misNum++
    }
    if misUpperChar {
        misNum++
    }

    dOne, dTwo, rep := 0, 0, 0
    for i := 2; i < len(s); i++ {
        if s[i] == s[i-1] && s[i-1] == s[i-2] {
            length := 3
            for i + 1 < len(s) && s[i+1] == s[i] {
                i++
                length++
            }
            if length % 3 == 0 {
                dOne++
            } else if length % 3 == 1 {
                dTwo++
            }

            rep += length / 3
        }
    }

    n := len(s)

    if n < 6 {
        return max(6 - n, misNum)
    } else if n <= 20 {
        return max(rep, misNum)
    } 

    del := n - 20
    rep -= min(del, dOne)
    if del > dOne {
        rep -= min((del - dOne) / 2, dTwo)
    }
    if del - dOne - 2 * dTwo > 0 {
        rep -= (del - dOne - 2 * dTwo) / 3
    }

    return del + max(rep, misNum)
}