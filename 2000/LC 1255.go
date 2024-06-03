func maxScoreWords(words []string, letters []byte, score []int) (ans int) {
    m := make([]int, 26)
    for _, c := range letters {
        m[int(c - 'a')]++
    }
    n := len(words)

    for i := 1; i < (1 << n); i++ {
        cnt := make([]int, 26)
        val := 0
        for j := range words {
            if i & (1 << j) != 0 {
                for _, c := range words[j] {
                    idx := int(c - 'a')
                    cnt[idx]++
                    val += score[idx]
                }
            }
        }
        f := true
        for i := range m {
            if m[i] < cnt[i] {
                f = false
                break
            }
        }
        if f {
            ans = max(ans, val)
        }
    }

    return 
}
