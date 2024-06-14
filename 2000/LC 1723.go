func minimumTimeRequired(jobs []int, k int) int {
    n := len(jobs)
    sort.Ints(jobs)
    sum := make([]int, 1 << n)
    for i := range jobs {
        for j, s := 0, 1 << i; j < s; j++ {
            sum[s|j] = sum[j] + jobs[i]
        }
    }

    f := make([]int, 1 << n)
    copy(f, sum)
    for i := 1; i < k; i++ {
        for j := (1 << n) - 1; j > 0; j-- {
            for s := j; s > 0; s = (s - 1) & j {
                f[j] = min(f[j], max(sum[s], f[s ^ j]))
            }
        }
    }

    return f[len(f)-1]
}