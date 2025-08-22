const N int = 26
func numSmallerByFrequency(queries []string, words []string) []int {
    ans := make([]int, len(queries))
    n := len(words)
    cnt := make([]int, n)

    calc := func(s string) (res int) {
        m := [N]int{}
        for _, c := range s {
            m[int(c - 'a')]++
        }
        for _, x := range m {
            if x != 0 {
                res = x
                break
            }
        }

        return 
    }

    for i, word := range words {
        cnt[i] = calc(word)
    }

    sort.Ints(cnt)
    for i, q := range queries {
        x := calc(q)
        pos := sort.SearchInts(cnt, x + 1)
        ans[i] = n - pos
    }

    return ans
}