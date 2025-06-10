func checkValidString(s string) bool {
    left, cnt := 0, 0
    for _, c := range s {
        if c == '(' {
            left++
        } else if c == ')' {
            if left > 0 {
                left--
            } else if cnt > 0 {
                cnt--
            } else {
                return false
            }
        } else {
            cnt++
        }
    }

    right, cnt := 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ')' {
            right++
        } else if s[i] == '(' {
            if right > 0 {
                right--
            } else if cnt > 0 {
                cnt--
            } else {
                return false
            }
        } else {
            cnt++
        }
    }
    return true
}