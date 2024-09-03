func minCharacters(a string, b string) int {
    cntA := make([]int, 26)
    cntB := make([]int, 26)
    for _, c := range a {
        x := int(c - 'a')
        cntA[x]++
    }
    for _, c := range b {
        x := int(c - 'a')
        cntB[x]++
    }

    m, n := len(a), len(b)
    ans := m + n
    sumA, sumB := cntA[0], cntB[0]
    for i := 1; i <= 25; i++ {
        ans = min(ans, m - sumA + sumB, n - sumB + sumA)
        sumA += cntA[i]
        sumB += cntB[i]
    }

    return min(ans, m - slices.Max(cntA) + n - slices.Max(cntB))
}


