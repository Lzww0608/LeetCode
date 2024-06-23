func calculate(s string) int {
    ans, n := 0, len(s)
    var cal func(idx int, b bool) int
    cal = func(idx int, b bool) int {
        if idx >= n {
            return 0
        }
        i, res := idx, 0
        f, t := true, 0
        for ; i < n && s[i] != ')'; i++ {
            if s[i] == '(' {
                i = cal(i+1, b == f)
                continue
            } else if s[i] == '+'{
                f = true
            } else if s[i] == '-' {
                f = false
            } else if unicode.IsDigit(rune(s[i])) {
                for i < n && unicode.IsDigit(rune(s[i])) {
                    t = t * 10 + int(s[i] - '0')
                    i++
                }
                i--
            } 
                if f {
                    res += t
                } else {
                    res -= t
                }
                t = 0
            
        }
        if f {
            res += t
        } else {
            res -= t
        }
        fmt.Println(b, res)
        if b {
            ans += res
        } else {
            ans -= res
        }
        return i
    }
    cal(0, true)
    return ans
}
