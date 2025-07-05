func minAddToMakeValid(s string) (ans int) {
    cur, cnt := 0, 0 
    for _, c := range s {
        if c == '(' {
            cnt++
        } else {
            cnt--
            if cnt < 0 {
                cur -= cnt 
                cnt = 0
            }
        }
    }

    ans += cur
    cur, cnt = 0, 0 
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ')' {
            cnt++
        } else {
            cnt--
            if cnt < 0 {
                cur -= cnt 
                cnt = 0
            }
        }
    }

    ans += cur
    return 
}