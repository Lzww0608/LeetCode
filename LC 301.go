func removeInvalidParentheses(s string) (ans []string) {
    l, r := 0, 0
    n := len(s)
    m := map[string]bool{}
    for i := range s {
        if s[i] == '(' {
            l++
        } else if s[i] == ')'{
            if l == 0 {
                r++
            } else {
                l--
            }
        }
    }

    check := func(s string) bool {
        cnt := 0
        for i := range s {
            if s[i] =='(' {
                cnt++
            } else if s[i] == ')' {
                cnt--
                if cnt < 0 {
                    return false
                }
            }
        }
        return cnt == 0
    }

    path := make(map[int]bool)
    var dfs func(int, int, int) 
    dfs = func(idx, l, r int) {
        if l == 0 && r == 0 {
            builder := strings.Builder{}
            for i := range s {
                if !path[i] {
                    builder.WriteByte(s[i])
                } 
            }
            if check(builder.String()) {
                m[builder.String()] = true
            }
            return
        }
        if l + r > n - idx {
            return
        }
        if s[idx] == '(' && l > 0 {
            path[idx] = true
            dfs(idx + 1, l - 1, r)
            delete(path, idx)
        } else if s[idx] == ')' && r > 0 {
            path[idx] = true
            dfs(idx + 1, l, r - 1)
            delete(path, idx)
        }
        dfs(idx + 1, l, r)
    }

    dfs(0, l, r)
    for k, _ := range m {
        ans = append(ans, k)
    }
    return
}