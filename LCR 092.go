func minFlipsMonoIncr(s string) int {
    a, b := 0, 0
    if s[0] == '0' {
        b = 1
    } else {
        a = 1
    }

    for i := 1; i < len(s); i++ {
        if s[i] == '0' {
            b = min(a, b) + 1
        } else {
            b = min(a, b)
            a++
        }
    }

    return min(a, b)
}