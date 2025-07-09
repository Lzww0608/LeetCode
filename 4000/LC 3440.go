func maxFreeTime(eventTime int, startTime []int, endTime []int) (ans int) {
    n := len(startTime)
    suf := make([]int, n + 2)
    pre := eventTime
    for i := n - 1; i >= 0; i-- {
        cur := pre - endTime[i]
        suf[i+1] = max(suf[i+2], cur)
        pre = startTime[i]
    }
    suf[0] = max(suf[1], pre)

    pre, cur := 0, 0
    for i := 0; i < n; i++ {
        d := endTime[i] - startTime[i]
        l, r := startTime[i] - pre, eventTime - endTime[i]
        if i + 1 < n {
            r = startTime[i+1] - endTime[i]
        }
        ans = max(ans, l + r)
        if cur >= d || suf[i+2] >= d {
            ans = max(ans, l + r + d)
        }
        cur = max(cur, l)
        pre = endTime[i]
    }

    return 
}