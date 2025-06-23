func oneEditAway(s string, t string) bool {
    if len(s) == len(t) {
        cnt := 0
        for i := range s {
            if s[i] != t[i] {
                cnt++
            }
        }

        return cnt <= 1
    } else if len(s) > len(t) {
        s, t = t, s
    }

    if len(s) + 1 != len(t) {
        return false
    }

    cnt := 0

    for i, j := 0, 0; j < len(t); {
        if i < len(s) {
            if s[i] == t[j] {
                i++
                j++
            } else {
                j++
                cnt++
            }
        } else {
            j++
        }

        if cnt > 1 {
            return false
        }
    }

    return true
}