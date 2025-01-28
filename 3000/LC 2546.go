func makeStringsEqual(s string, target string) bool {
    // 0 0: 0 0
    // 1 0: 1 1
    // 1 1: 1 0
    cntS, cntT := 0, 0
    for i := range s {
        if s[i] == target[i] && s[i] == '1' {
            return true
        }

        if s[i] == '1' {
            cntS++
        } 
        if target[i] == '1' {
            cntT++
        }
    }

    return cntS == 0 && cntT == 0 || cntS > 0 && cntT > 0
}