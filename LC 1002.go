func commonChars(words []string) (ans []string) {
    cnt := make([]int, 26)
    for i := range cnt {
        cnt[i] = 100
    }
    for _, s := range words {
        tmp := make([]int, 26)
        for i := range s {
            tmp[int(s[i] - 'a')]++
        }
        for i := range tmp {
            cnt[i] = min(cnt[i], tmp[i])
        }
    }

    for i := range cnt {
        for cnt[i] > 0 {
            cnt[i]--
            ans = append(ans, string(byte('a' + i)))
        }
    }

    return 
}