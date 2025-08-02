func maxScore(cardPoints []int, k int) (ans int) {
    n := len(cardPoints)
    suf := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        suf[i] = suf[i + 1] + cardPoints[i]
    }

    pre := 0
    for i := 0; i < min(k + 1, n); i++ {
        x := cardPoints[i]
        ans = max(ans, pre + suf[n - k + i])
        pre += x
    }

    return 
}