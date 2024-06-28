func wordSubsets(words1 []string, words2 []string) []string {
    ans := []string{}

    m := make([]int, 26)
    for _, s := range words2 {
        mm := make([]int, 26)
        for _, c := range s {
            mm[c-'a']++
        }
        for i, x := range mm {
            m[i] = max(m[i], x)
        }
    }

    for _, s := range words1 {
        mm := make([]int, 26)
        for _, c := range s {
            mm[c-'a']++
        }
        f := true
        for i := range mm {
            if mm[i] < m[i] {
                f = false
                break
            }
        }
        if f {
            ans = append(ans, s)
        }
    }
    return ans
}
