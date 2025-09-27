const N int = 26
func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
    ans := 0
    pre := [N + 1]int{}
    suf := [N + 1]int{}
    for i := 0; i < N; i++ {
        pre[i + 1] = pre[i] + nextCost[i]
        j := N - i - 1
        suf[j] = suf[j + 1] + previousCost[j]
    }
    for i := range s {
        x := int(s[i] - 'a')
        y := int(t[i] - 'a')
        if x == y {
            continue
        }
        var forwardCost, bacwordCost int 
        if x < y {
            forwardCost = pre[y] - pre[x]
            bacwordCost = suf[0] - suf[x + 1] + suf[y + 1]
        } else {
            forwardCost = pre[N] - pre[x] + pre[y]
            bacwordCost = suf[y + 1] - suf[x + 1]
        }
        ans += min(forwardCost, bacwordCost)
    }

    return int64(ans)
}