func jobScheduling(startTime []int, endTime []int, profit []int) int {
    n := len(startTime)
    jobs := make([][3]int, n)
    for i := range jobs {
        a, b, c := startTime[i], endTime[i], profit[i]
        jobs[i][0], jobs[i][1], jobs[i][2] = a, b, c
    }
    sort.Slice(jobs, func(i, j int) bool {
        if jobs[i][1] == jobs[j][1] {
            return jobs[i][0] < jobs[j][0]
        }
        return jobs[i][1] < jobs[j][1]
    })
    f := make([]int, n + 1)
    for i := range jobs {
        a, c := jobs[i][0], jobs[i][2]
        f[i+1] = max(f[i], c)
        l, r := 0, i
        for l < r {
            mid := l + ((r - l) >> 1)
            if jobs[mid][1] <= a {
                l = mid + 1
            } else {
                r = mid
            }
        }
        f[i+1] = max(f[i+1], f[l] + c)
    }

    return f[n]
}
