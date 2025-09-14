func doesAliceWin(s string) bool {
    cnt := 0
    m := map[byte]bool {
        'a': true, 
        'e': true,
        'i': true, 
        'o': true,
        'u': true,
    }

    for i := range s {
        if m[s[i]] {
            cnt++
        }
    }

    if cnt == 0 {
        return false
    }

    return true
}