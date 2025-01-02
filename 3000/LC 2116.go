func canBeValid(s string, locked string) bool {
    n := len(s)
    if n & 1 == 1 {
        return false
    }

    cnt := 0
    st := []byte{}
    for i := range s {
        if locked[i] == '0' {
            cnt++
        } else {
            if s[i] == '(' {
                st = append(st, s[i])
            } else {
                if len(st) > 0 {
                    st = st[:len(st)-1]
                } else {
                    if cnt > 0 {
                        cnt--
                    } else {
                        return false
                    }
                }
            }
        }
    }

    cnt = 0
    st = []byte{}
    for i := n - 1; i >= 0; i-- {
        if locked[i] == '0' {
            cnt++
        } else {
            if s[i] == ')' {
                st = append(st, s[i])
            } else {
                if len(st) > 0 {
                    st = st[:len(st)-1]
                } else {
                    if cnt > 0 {
                        cnt--
                    } else {
                        return false
                    }
                }
            }
        }
    }

    return true
}