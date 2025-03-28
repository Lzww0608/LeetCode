func minimizedStringLength(s string) int {
    cnt := make(map[byte]int)
    for i := range s {
        cnt[s[i]]++
    }

    return len(cnt)
}