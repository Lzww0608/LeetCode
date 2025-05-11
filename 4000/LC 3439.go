func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) (ans int) {
    n := len(startTime)
    get := func(x int) int {
        if x == 0 {
            return startTime[0]
        } else if x == n {
            return eventTime - endTime[n-1]
        }

        return startTime[x] - endTime[x-1]
    }

    cur := 0
    for i := 0; i <= n; i++ {
        cur += get(i)
        ans = max(ans, cur)
        if i >= k {
            cur -= get(i-k)
        } 
    }

    return 
}