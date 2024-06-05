func resultArray(nums []int) []int {
    //n := len(nums)
    arr := append([]int(nil), nums...)
    sort.Ints(arr)
    arr = slices.Compact(arr)
    m := len(arr)

    update := func(f []int, idx, val int) {
        for i := idx; i < len(f); i += i & (-i) {
            f[i] += val
        }
    }

    query := func(f []int, idx int) (res int) {
        for i := idx; i > 0; i -= i & (-i) {
            res += f[i]
        }
        return
    }

    a := nums[:1]
    b := []int{nums[1]}

    f := make([]int, m + 1)
    g := make([]int, m + 1)
    update(f, sort.SearchInts(arr, nums[0]) + 1, 1)
    update(g, sort.SearchInts(arr, nums[1]) + 1, 1)

    for _, x := range nums[2:] {
        pos := sort.SearchInts(arr, x) + 1
        sel_a := len(a) - query(f, pos)
        sel_b := len(b) - query(g, pos)

        if sel_a > sel_b || (sel_a == sel_b && len(a) <= len(b)) {
            a = append(a, x)
            update(f, pos, 1)
        } else {
            b = append(b, x)
            update(g, pos, 1)
        }
    }

    return append(a, b...)
}