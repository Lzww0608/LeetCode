func secondHighest(s string) int {
    a, b := -1, -1
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            x := int(s[i] - '0')
            if x > a {
                a, b = x, a
            } else if x != a && x > b {
                b = x 
            }
        }
    }

    return b
}