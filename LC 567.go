func checkInclusion(s1 string, s2 string) bool {
    m := map[int]int{}
    for _, c := range s1 {
        m[int(c-'a')]++
    }
    l, r, n := 0, 0, len(s2)
    for r < n {
        x := int(s2[r]-'a')
        m[x]--
        for l < len(s2) && m[x] < 0 {
            m[int(s2[l]-'a')]++
            l++
        }
        if r - l + 1 == len(s1) {
            return true
        }
        r++
    }
    return false
    
}