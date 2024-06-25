func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := map[byte]int{}
    
    for i := range s {
        m[s[i]]++
        m[t[i]]--
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}
