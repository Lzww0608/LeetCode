func countRatioSubarrays(nums []int, a int, b int) int64 {
    // x * b <= a * y
    // a * y - x * b >= 0, x: even, y: odd
    for i, x := range nums {
        if x & 1 == 1 {
            nums[i] = a
        } else {
            nums[i] = -b
        }
    }
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i + 1] = pre[i] + x
    }
    sum := append([]int(nil), pre...)
    sort.Ints(sum)
    sum = slices.Compact(sum)

    f := make([]int, len(sum) + 1)
    update := func(i, x int) {
        for i < len(f) {
            f[i] += x 
            i += i & (-i)
        }
    }

    query := func(i int) (res int) {
        for i > 0 {
            res += f[i]
            i -= i & (-i)
        }

        return
    }

    ans := 0
    for _, x := range pre {
        y := sort.SearchInts(sum, x) + 1
        ans += query(y)
        update(y, 1)
    }

    return int64(ans)
}