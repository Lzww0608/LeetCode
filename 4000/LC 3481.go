func applySubstitutions(replacements [][]string, text string) string {
    m := make(map[string]string)
    for _, v := range replacements {
        m[v[0]] = v[1]
    }

    ans := text
    for {
        f := true 
        tmp := []byte{}
        for i := 0; i < len(ans); i++ {
            if ans[i] != '%' {
                tmp = append(tmp, ans[i])
            } else {
                j := i + 1
                for j < len(ans) && ans[j] != '%' {
                    j++
                }
                if j < len(ans) {
                    f = false
                    tmp = append(tmp, m[ans[i + 1: j]]...)
                    i = j
                }
            }
        }
        if f {
            break
        }
        ans = string(tmp)
    }

    return ans
}