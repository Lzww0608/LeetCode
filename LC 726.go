func countOfAtoms(formula string) string {
    m := map[string]int{}
    st := []int{}
    var d string
    builder := strings.Builder{}
    cnt := 1
    for i := len(formula) - 1; i >= 0; i-- {
        c := formula[i]
        if c >= '0' && c <= '9' {
            d = string(c) + d
        } else if c == ')' {
            v := 1
            if d != "" {
                v, _ = strconv.Atoi(d)
            }
            d = ""
            cnt *= v
            st = append(st, v)
        } else if c == '(' {
            cnt /= st[len(st)-1]
            st = st[:len(st)-1]
        } else {
            builder.WriteByte(c)
            if c >= 'A' && c <= 'Z' {
                s := []byte(builder.String())
                for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
                    s[i], s[j] = s[j], s[i]
                }
                if d == "" {
                    m[string(s)] += cnt
                } else {
                    t, _ := strconv.Atoi(d)
                    m[string(s)] += t * cnt
                }
                builder.Reset()
                d = ""
            }
        }
    }
    builder.Reset()
    type pair struct {
        s string
        x int
    }
    a := make([]pair, 0, len(m))
    for k, v := range m {
        a = append(a, pair{k, v})
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i].s < a[j].s
    })
    for _, v := range a {
        builder.WriteString(v.s)
        if v.x > 1 {
            builder.WriteString(strconv.Itoa(v.x))
        }
    }

    return builder.String()
}