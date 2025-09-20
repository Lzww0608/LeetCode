const N int = 26
func removeAnagrams(words []string) []string {
    a := []int{}

    pre := [N]int{}
    for i := 0; i < len(words); i++ {
        cnt := [N]int{}
        for j := range words[i] {
            x := int(words[i][j] - 'a')
            cnt[x]++
        }
        if cnt != pre {
            pre = cnt
            a = append(a, i)
        }
    }

    ans := make([]string, len(a))
    for i := range a {
        ans[i] = words[a[i]]
    }

    return ans
}