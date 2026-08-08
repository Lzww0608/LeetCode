func countTasks(tasks []int, shifts []int) []int {
    m, n := len(tasks), len(shifts)
    ans := make([]int, n)
    pre := make([]int, m + 1)
    for i, x := range tasks {
        pre[i + 1] = pre[i] + x
    }

    j, add := 0, 0
    for i, x := range shifts {
        cur := pre[j] + add + x
        if cur >= pre[m] {
            j, add = 0, 0
            continue
        }
        j = sort.SearchInts(pre, cur)
        if pre[j] != cur {
            add = cur - pre[j]
            ans[i] = m - j + 1
        } else {
            add = 0
            ans[i] = m - j
        }
        
    }

    return ans
}