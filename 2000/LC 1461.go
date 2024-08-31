func hasAllCodes(s string, k int) bool {
    m := map[string]struct{}{}
    for i := 0; i + k <= len(s); i++ {
        m[s[i:i+k]] = struct{}{}
    }

    if len(m) == 1 << k {
        return true
    }

    return false
    
}