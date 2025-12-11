func minInversionCount(nums []int, k int) int64 {
    n := len(nums)
    tmp := make([]int, n)
    copy(tmp, nums)
    m := make(map[int]int)
    sort.Ints(tmp)
    id := 1
    for i := range n {
        if i == 0 || tmp[i] != tmp[i - 1] {
            m[tmp[i]] = id
            id++
        }
    }

    f := make([]int, id + 1)
    update := func(i, val int) {
        for ; i <= id; i += i & (-i) {
            f[i] += val
        }
    }

    query := func(i int) (res int) {
        for ; i > 0; i -= i & (-i) {
            res += f[i]
        }

        return
    }

    ans := 0
    for i := range k {
        ans += i - query(m[nums[i]])
        update(m[nums[i]], 1)
    }

    cur := ans
    for i := k; i < n; i++ {
        update(m[nums[i - k]], -1)
        cur -= query(m[nums[i - k]] - 1)
        cur += k - 1 - query(m[nums[i]])
        update(m[nums[i]], 1)
        ans = min(ans, cur)
    }

    return int64(ans)
}