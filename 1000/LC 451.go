func frequencySort(s string) string {
    cnt := make(map[byte]int)
    for i := range s {
        cnt[s[i]]++
    }    

    ans := []byte(s)
    sort.Slice(ans, func(i, j int) bool {
        if cnt[ans[i]] == cnt[ans[j]] {
            return ans[i] < ans[j]
        }
        return cnt[ans[i]] > cnt[ans[j]]
    })

    return string(ans)
}