func longestValidSubstring(word string, forbidden []string) (ans int) {
    m := map[string]bool{}
    for _, v := range forbidden {
        m[v] = true
    }

    s := make([]byte, 0, len(word))
    for i := range word {
        s = append(s, word[i])
        for j := 1; j <= min(len(s), 10); j++ {
            if m[string(s[len(s)-j:])] {
                s = s[len(s)-j+1:]
                break
            }
        }
        ans = max(ans, len(s))
    }

    return
}