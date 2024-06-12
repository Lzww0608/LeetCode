func lengthOfLongestSubstring(s string) (ans int) {
    m := map[byte]int{}
    j := 0
    for i := range s {
        m[s[i]]++
        if m[s[i]] == 1 {
            ans = max(ans, len(m))
        } else {
            for m[s[i]] > 1 {
                m[s[j]]--
                if m[s[j]] == 0 {
                    delete(m, s[j])
                }
                j++
            }
        }
    }

    return 
}
