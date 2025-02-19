func reformatNumber(number string) string {
    s := []byte{}
    for i := range number {
        if number[i] >= '0' && number[i] <= '9' {
            s = append(s, number[i])
        }
    }

    n := len(s)
    ans := []byte{}
    for i := 0; i < len(s); {
        if n > 4 {
            ans = append(ans, s[i:i+3]...)
            ans = append(ans, '-')
            i += 3
            n -= 3
        } else if n == 4 {
            ans = append(ans, s[i:i+2]...)
            ans = append(ans, '-')
            ans = append(ans, s[i+2:]...)
            break
        } else {
            ans = append(ans, s[i:]...)
            break
        }
    }

    return string(ans)
}