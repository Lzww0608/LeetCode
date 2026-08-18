const N int = 26
func maximumGap(skill string, station string) (ans int) {
    m, n := len(skill), len(station)
    if m == 1 {
        return
    }

    suf := make([]int, n)
    for i, j := n - 1, m - 1; i >= 0 && j >= 0; i-- {
        x := int(station[i] - 'a')
        suf[i] = j
        if j >= 0 && x == int(skill[j] - 'a') {
            j--
        }
    }

    j := 0
    for i := range n {
        if j > 0 {
            k := sort.SearchInts(suf, j)
            ans = max(ans, k - i + 1)
        }
        
        x := int(skill[j] - 'a')
        y := int(station[i] - 'a')
        if x == y {
            j++
            if j >= m {
                break
            }
        }
    }

    return
}