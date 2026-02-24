func kthCharacter(s string, k int64) byte {
    var cnt int64 = 0
    for i := range s {
        if s[i] == ' ' {
            cnt = 1
        } else {
            if i == 0 || s[i - 1] != ' ' {
                cnt++
            } else {
                cnt = 1
            }
        }

        if k < cnt {
            return s[i]
        }
        k -= cnt
    }

    return ' '
}