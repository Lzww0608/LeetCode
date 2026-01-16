func residuePrefixes(s string) (ans int) {
    cnt := make(map[byte]bool)
    for i := range s {
        cnt[s[i]] = true 
        if len(cnt) >= 3 {
            break
        }
        if len(cnt) == (i + 1) % 3 {
            ans++
        }
    }

    return
}