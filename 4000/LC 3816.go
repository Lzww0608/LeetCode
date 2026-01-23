func lexSmallestAfterDeletion(s string) string {
    cnt := make(map[byte]int)
    for i := range s {
        cnt[s[i]]++
    }

    st := []byte{}
    for i := range s {
        for len(st) > 0 && st[len(st) - 1] > s[i] && cnt[st[len(st) - 1]] > 1 {
            cnt[st[len(st) - 1]]--
            st = st[:len(st) - 1]
        }
        st = append(st, s[i])
    }

    for len(st) > 0 && cnt[st[len(st) - 1]] > 1 {
        cnt[st[len(st) - 1]]--
        st = st[:len(st) - 1]
    }

    return string(st)
}