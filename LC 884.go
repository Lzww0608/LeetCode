func uncommonFromSentences(s1 string, s2 string) []string {
    ans := []string{}
    a := strings.Split(s1, " ")
    b := strings.Split(s2, " ")
    m := map[string]int{}
    for _, ss := range a {
        m[ss]++
    }
    for _, ss := range b {
        m[ss]++
    }
    for k, v := range m {
        if v == 1 {
            ans = append(ans, k)
        }
    }
    return ans
}