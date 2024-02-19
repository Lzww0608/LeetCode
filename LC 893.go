func numSpecialEquivGroups(words []string) int {
    m := map[string]int{}
    ans := 0
    for _, word := range words {
        odd := strings.Builder{}
        even := strings.Builder{}
        for i, c := range word {
            if (i & 1) == 1 {
                odd.WriteRune(c)
            } else {
                even.WriteRune(c)
            }
        }
        a, b := []rune(odd.String()), []rune(even.String())
        sort.Slice(a, func(i, j int) bool {
            return a[i] < a[j]
        })
        sort.Slice(b, func(i, j int) bool {
            return b[i] < b[j]
        })
        s := string(a) + string(b)
        if _, ok := m[s]; !ok {
            ans++
            m[s] = 1
        }
    }
    return ans
}