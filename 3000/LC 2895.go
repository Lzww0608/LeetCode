func minProcessingTime(processorTime []int, tasks []int) (ans int) {
    sort.Ints(processorTime)
    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i] > tasks[j]
    })

    for i, x := range tasks {
        ans = max(ans, processorTime[i / 4] + x)
    }

    return
}