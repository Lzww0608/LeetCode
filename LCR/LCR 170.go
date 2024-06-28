func reversePairs(record []int) (ans int) {
    n := len(record)
    m := map[int]int{}
    tmp := make([]int, n)
    copy(tmp, record)
    sort.Ints(tmp)

    for i, x := range tmp {
        m[x] = i + 1
    }

    f := make([]int, n + 1)
    
    update := func(idx, val int) {
        for i := idx; i < len(f); i += i & (-i) {
            f[i] += val
        }
    }

    query := func(idx int) (res int) {
        for i := idx; i > 0; i -= i & (-i) {
            res += f[i]
        }
        return
    }

    for i := n - 1; i >= 0; i-- {
        ans += query(m[record[i]] - 1)
        update(m[record[i]], 1)
    }

    return
}
