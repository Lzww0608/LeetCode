func isIsomorphic(s string, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}
    for i := range s {
        a, ok1 := m1[s[i]]
        b, ok2 := m2[t[i]] 
        if !ok1 && !ok2 {
            m1[s[i]] = t[i]
            m2[t[i]] = s[i]
        } else if ok1 && ok2 && a == t[i] && b == s[i] {
            continue
        } else {
            return false
        }
    }

    return true
}
