func numberOfSubstrings(s string) int64 {
    ans := 0
    cnt := [26]int{}
    for _, c := range s {
        x := c - 'a'
        cnt[x]++
        ans += cnt[x]
    }

    return int64(ans)
}