func minimumTime(jobs []int, workers []int) (ans int) {
    n := len(jobs)
    sort.Ints(jobs)
    sort.Ints(workers)
    for i := 0; i < n; i++ {
        ans = max(ans, (jobs[i] + workers[i] - 1) / workers[i])
    }

    return
}