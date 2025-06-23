func findSubstring(s string, words []string) (ans []int) {
     m :=  len(words[0])
    exists := make(map[string]int)
    for _, word := range words {
        exists[word]++
    }

    for t := 0; t < m; t++ {
        vis := make(map[string]int)
        cnt := 0
        for i, l := t, t; i + m <= len(s); i += m {
            if v, ok := exists[s[i:i+m]]; ok {
                vis[s[i:i+m]]++
                if v == vis[s[i:i+m]] {
                    cnt++
                } 
                for v < vis[s[i:i+m]] {
                    if vis[s[l:l+m]] == exists[s[l:l+m]] {
                        cnt--
                    }
                    vis[s[l:l+m]]--
                    l += m
                }
                if cnt == len(exists) {
                    ans = append(ans, l)
                    vis[s[l:l+m]]--
                    l += m
                    cnt--
                }
            } else {
                vis = make(map[string]int)
                cnt = 0
                l = i + m
            }
        }
    }

    return 
}