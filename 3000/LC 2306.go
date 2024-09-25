func distinctNames(ideas []string) int64 {
    s := make([]map[string]bool, 26)
    for i := range s {
        s[i] = make(map[string]bool)
    }
    ans := 0

    for _, idea := range ideas {
        t := (idea[0] & 31) - 1 
        s[t][idea[1:]] = true
    }


    for i := 0; i < 26; i++ {
        for j := i; j < 26; j++ {
            a, b := len(s[i]), len(s[j])
            for t := range s[i] {
                if s[j][t] {
                    a--
                }
            }
            for t := range s[j] {
                if s[i][t] {
                    b--
                }
            }

            ans += a * b * 2
        }
    }

    return int64(ans)
}