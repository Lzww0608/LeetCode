const N int = 10
func findValidPair(s string) string {
    cnt := [N]int{}
    for i := range s {
        x := int(s[i] - '0')
        cnt[x]++
    }

    for i := range len(s) - 1 {
        a, b := int(s[i] - '0'), int(s[i + 1] - '0')
        if a != b && cnt[a] == a && cnt[b] == b {
            return s[i: i + 2]
        }
    }
    
    return ""
}