func countWordOccurrences(chunks []string, queries []string) []int {
    s := strings.Join(chunks, "")
    n := len(s)

    cnt := make(map[string]int)
    for i := 0; i < n; i++ {
        if s[i] == ' ' || s[i] == '-' {
            continue
        }
        j := i 
        cur := []byte{}
        for j < n && s[j] != ' ' {
            if s[j] == '-' && (j + 1 == n || s[j + 1] == ' ' || s[j + 1] == '-') {
                j++
                break
            }

            cur = append(cur, s[j])
            j++
        } 
        cnt[string(cur)]++
        i = j
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        ans[i] = cnt[q]
    }

    return ans
}