func countOperationsToEmptyArray(nums []int) int64 {
    n := len(nums)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return nums[id[i]] < nums[id[j]]
    })

    f := make([]int, n + 1)
    update := func(i int) {
        for ; i < len(f); i += i & (-i) {
            f[i]++
        }
    }

    pre := func(i int) (res int) {
        for ; i > 0; i -= i & (-i) {
            res += f[i]
        }
        return
    }

    query := func(l, r int) int {
        if r < l {
            return 0
        }

        return pre(r) - pre(l - 1) 
    }

    ans := n
    last := 1
    for _, i := range id {
        i++
        if i >= last {
            ans += i - last - query(last, i)
        } else {
            ans += (n - last - query(last, n)) + (i - query(1, i))
        }
        update(i)
        last = i
    }

    return int64(ans)
}