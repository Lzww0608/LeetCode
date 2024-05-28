func groupAnagrams(strs []string) (ans [][]string) {
    m := map[[26]int][]string{}
    for _, s := range strs {
        cnt := [26]int{}
        for _, b := range s {
            cnt[int(b-'a')]++
        }
        m[cnt] = append(m[cnt], s)
    }
    for _, v := range m {
        ans = append(ans, v)
    }
    return
}