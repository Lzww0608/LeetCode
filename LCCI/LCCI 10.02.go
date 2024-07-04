func groupAnagrams(strs []string) (ans [][]string) {
    m := map[[26]int]int{}
    for _, s := range strs {
        cnt := [26]int{}
        for _, c := range s {
            x := int(c - 'a')
            cnt[x]++
        }
        if _, ok := m[cnt]; !ok {
            m[cnt] = len(ans)
            ans = append(ans, []string{s})
        } else {
            i := m[cnt]
            ans[i] = append(ans[i], s)
        }
    }

    return
}