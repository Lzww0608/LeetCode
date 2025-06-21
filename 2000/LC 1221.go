func balancedStringSplit(s string) (ans int) {
    cnt := 0
    for i := range s {
        if s[i] == 'L' {
            cnt++
        } else {
            cnt--
        }
        if cnt == 0 {
            ans++
        }
    }

    return
}