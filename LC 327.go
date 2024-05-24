func countRangeSum(nums []int, lower int, upper int) (ans int) {
    n := len(nums)
    sum := make([]int, n + 1)
    for i, x := range nums {
        sum[i+1] = sum[i] + x
    }

    a := make([]int, 0, 3 * n + 1)
    for _, x := range sum {
        a = append(a, x)
        a = append(a, x - lower)
        a = append(a, x - upper)
    }
    sort.Ints(a)

    m := map[int]int{}
    i := 0
    for _, x := range a {
        m[x] = i 
        i++
    }

    f := make([]int, len(a) + 1)
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

    for i := range sum {
        l, r := m[sum[i]-upper], m[sum[i]-lower]
        ans += query(r + 1) - query(l)
        update(m[sum[i]] + 1, 1)
    }

    return
}