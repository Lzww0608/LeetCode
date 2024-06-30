func reversePairs(nums []int) (ans int) {
    a := append([]int(nil), nums...)
    for _, x := range nums {
        a = append(a, x * 2)
    }
    sort.Ints(a)

    m := map[int]int{}
    for i, x := range a {
        m[x] = i + 1
    }

    sum := make([]int, len(a) + 1)

    update := func(idx, val int) {
        for i := idx; i < len(sum); i += i & (-i) {
            sum[i] += val
        }
    }

    query := func(idx int) (res int) {
        for i := idx; i > 0; i -= i & (-i) {
            res += sum[i]
        }
        return
    }

    for i := len(nums) - 1; i >= 0; i-- {
        ans += query(m[nums[i]] - 1)
        update(m[nums[i] * 2], 1)
    }

    return
}
