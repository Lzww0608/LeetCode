func maximumSetSize(a []int, b []int) int {
    n := len(a)
    cntA, cntB := make(map[int]bool), make(map[int]bool)
    
    for i := 0; i < n; i++ {
        cntA[a[i]] = true
        cntB[b[i]] = true
    }

    cnt := len(cntA) + len(cntB)
    for k := range cntA {
        if cntB[k] {
            cnt--
        }
    }
    cA := min(n / 2, len(cntA))
    cB := min(n / 2, len(cntB))

    return min(cnt, cA + cB)
}