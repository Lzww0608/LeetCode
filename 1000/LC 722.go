func removeComments(source []string) (ans []string) {
    cur := []byte{}
    pre := 0
    for _, s := range source {
        n := len(s)
        for i := 0; i < n; i++ {
            if s[i] == '/' && i < n - 1 && s[i+1] == '/' && pre == 0 {
                if len(cur) != 0 {
                    ans = append(ans, string(cur))
                    cur = []byte{}
                }
                break
            }

            if s[i] == '/' && i < n - 1 && s[i+1] == '*' && pre == 0 {
                pre = 1
                i++
                continue
            } 
            if s[i] == '*' && i < n - 1 && s[i+1] == '/' && pre == 1 {
                pre = 0
                i++
                continue
            }

            if pre != 1 {
                cur = append(cur, s[i])
            }
        }
        if pre != 1 && len(cur) > 0 {
            ans = append(ans, string(cur))
            cur = []byte{}
        }
    }

    return
}