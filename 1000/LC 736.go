func evaluate(s string) int {
    i, n := 0, len(s)
    parseVar := func() string {
        j := i
        for i < n && s[i] != ' ' && s[i] != ')' {
            i++
        }

        return s[j:i]
    }

    parseInt := func() int {
        sign, x := 1, 0
        if s[i] == '-' {
            sign = -1
            i++
        }

        for i < n && unicode.IsDigit(rune(s[i])) {
            x = x * 10 + int(s[i] - '0')
            i++
        }

        return sign * x
    }

    m := map[string][]int{}
    var dfs func() int
    dfs = func() (res int) {
        if s[i] != '(' {
            if unicode.IsLower(rune(s[i])) {
                tmp := m[parseVar()]
                return tmp[len(tmp)-1]
            }
            return parseInt()
        }
        i++
        
        if s[i] == 'l' {
            i += 4
            cur := []string{}
            for {
                if !unicode.IsLower(rune(s[i])) {
                    res = dfs()
                    break
                }

                tmp := parseVar()
                if s[i] == ')' {
                    t := m[tmp]
                    res = t[len(t)-1]
                    break
                }

                cur = append(cur, tmp)
                i++
                m[tmp] = append(m[tmp], dfs())
                i++
            }
            for _, v := range cur {
                m[v] = m[v][:len(m[v])-1]
            }
        } else if s[i] == 'a' {
            i += 4
            e1 := dfs()
            i++
            e2 := dfs()
            res = e1 + e2
        } else {
            i += 5
            e1 := dfs()
            i++
            e2 := dfs()
            res = e1 * e2
        }
        i++
        return
    }

    return dfs()
}