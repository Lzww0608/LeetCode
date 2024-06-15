func minimumEffort(tasks [][]int) (ans int) {
    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i][1] - tasks[i][0] < tasks[j][1] - tasks[j][0]
    })
    for _, v := range tasks {
        ans = max(ans + v[0], v[1])
    }
    return 
}
