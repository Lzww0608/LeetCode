func numSplits(s string) (ans int) {
    cnt := map[byte]int{}
    for i := range s {
        cnt[s[i]]++
    }

    cur := map[byte]int{}
    for i := range s {
        cur[s[i]]++
        if cnt[s[i]]--; cnt[s[i]] == 0 {
            delete(cnt, s[i])
        }
        if len(cur) == len(cnt) {
            ans++
        }
    }

    return
}