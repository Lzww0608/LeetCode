const N int = 26
func numMatchingSubseq(s string, words []string) (ans int) {
    n := len(s)
    pos := [N]int{}
    for i := range N {
        pos[i] = n + 1
    }
    suf := make([][N]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        suf[i + 1] = pos 
        c := int(s[i] - 'a')
        pos[c] = i + 1
    }
    suf[0] = pos

    for _, word := range words {
        i := 0
        for _, c := range word {
            x := int(c - 'a')
            i = suf[i][x]
            if i > n {
                break
            }
        }

        if i <= n {
            ans++
        }
    }

    return
}