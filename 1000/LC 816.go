func ambiguousCoordinates(s string) (ans []string) {
    s = s[1:len(s)-1]
    n := len(s)

    get := func(s string) (res []string) {
        if s[0] != '0' || s == "0" {
            res = append(res, s)
        }
        for i := 1; i < len(s); i++ {
            if (i > 1 && s[0] == '0') || s[len(s)-1] == '0' {
                break
            }
            res = append(res, s[:i] + "." + s[i:])
        }
        return
    }

    for i := 1; i < n; i++ {
        l := get(s[:i])
        r := get(s[i:])
        for _, p := range l {
            for _, q := range r {
                ans = append(ans, "(" + p + ", " + q + ")")
            }
        }
    }

    return
}