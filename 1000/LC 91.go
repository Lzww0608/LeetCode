func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }

    a, b := 1, 1
    for i := 1; i < len(s); i++ {
        c := 0
        if s[i] != '0' {
            c += b
        }
        t := int(s[i-1] - '0') * 10 + int(s[i] - '0')
        if s[i-1] != '0' && t <= 26 {
            c += a
        }
        a, b = b, c
    }

    return b
}
