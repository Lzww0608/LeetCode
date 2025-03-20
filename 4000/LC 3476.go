func maxProfit(workers []int, tasks [][]int) int64 {
    ans, mx := 0, 0
    sort.Ints(workers)
    sort.Slice(tasks, func(i, j int) bool {
        if tasks[i][0] == tasks[j][0] {
            return tasks[i][1] > tasks[j][1]
        }
        return tasks[i][0] < tasks[j][0]
    })

    n, m := len(workers), len(tasks)
    for i, j := 0, 0; j < m; j++ {
        if i < n {
            for i < n && workers[i] < tasks[j][0] {
                i++
            }
            if i < n && workers[i] == tasks[j][0] {
                ans += tasks[j][1]
                i++
                continue
            } 
        } 
        mx = max(mx, tasks[j][1])
    }

    return int64(ans + mx)
}