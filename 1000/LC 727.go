func minWindow(s string, t string) (ans string) {
    m := len(s)

    pos := [26]int{}
    for i := range pos {
        pos[i] = m 
    }

    next := make([][26]int, m)
    for i := m - 1; i >= 0; i-- {
        next[i] = pos 
        pos[s[i]-'a'] = i 
    }

    for i := range s {
        if s[i] == t[0] {
            j := i 
            for _, c := range t[1:] {
                j = next[j][c-'a']
                if j == m {
                    return
                }
            }

            w := s[i:j+1]
            if ans == "" || len(w) < len(ans) {
                ans = w
            }
        }
    }
    
    return 
}