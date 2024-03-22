func areAlmostEqual(s1 string, s2 string) bool {
    var a, b byte
    cnt := 0
    for i := range s1 {
        if s1[i] != s2[i] {
            cnt++
            if cnt > 2 {
                return false
            } else if cnt == 1 {
                a, b = s1[i], s2[i]
            } else {
                if a != s2[i] || b != s1[i] {
                    return false
                }
            }
        }
    }

    return cnt != 1
}