func maxTotal(value []int, limit []int) int64 {
    n := len(value)
    id := make([]int, n)
    mx := 0
    for i := range n {
        id[i] = i
        mx = max(mx, limit[i])
    }

    sort.Slice(id, func(i, j int) bool {
        if limit[id[i]] == limit[id[j]] {
            return value[id[i]] > value[id[j]]
        }
        return limit[id[i]] < limit[id[j]]
    })

    ans := 0

    j := 0
    for i := 1; i <= mx && j < n; i++ {
        cnt := 0
        for j < n && limit[id[j]] == i {
            cnt++
            if cnt <= i {
                ans += value[id[j]]
            }
            j++
        }
    }

    return int64(ans)
}