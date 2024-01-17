func maximumNumberOfStringPairs(words []string) int {
    ans := 0
    m := map[string]int{}
    for _, s := range words {
        t := string(s[1]) + string(s[0])
        if m[t] > 0 {
            m[t]--
            ans++
        } else {
            m[s]++
        }
    }
    return ans;
}