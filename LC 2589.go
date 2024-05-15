func findMinimumTime(tasks [][]int) int {

    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i][1] < tasks[j][1]
    })
    st := [][]int{{-1, -1, 0}}
    for _, v := range tasks {
        l, r, d := v[0], v[1], v[2]
        pos := sort.Search(len(st), func(i int) bool {
            return st[i][0] >= l
        })
        d -= st[len(st)-1][2] - st[pos-1][2]
        if l <= st[pos-1][1] {
            d -= st[pos-1][1] - l + 1
        }
        if d <= 0 {
            continue
        }
        for r - st[len(st)-1][1] <= d {
            d += st[len(st)-1][1] - st[len(st)-1][0] + 1
            st = st[:len(st)-1]
        }
        st = append(st, []int{r - d + 1, r, st[len(st)-1][2] + d})
    }
    return st[len(st)-1][2]
}