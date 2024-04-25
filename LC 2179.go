func goodTriplets(nums1 []int, nums2 []int) (ans int64) {
    n := len(nums1)
    pos := make([]int, n)
    for i := range pos {
        pos[nums2[i]] = i + 1
    }

    a := make([]int, n + 1)

    update := func(idx, val int) {
        for i := idx; i <= n; i += i & (-i) {
            a[i] += val
        }
    }

    query := func(idx int) int {
        res := 0
        for i := idx; i > 0; i -= i & (-i) {
            res += a[i]
        }

        return res
    }

    for _, x := range nums1 {
        idx := pos[x]
        
        l := query(idx)
        r := n - idx - (query(n) - query(idx))

        ans += int64(l) * int64(r)

        update(idx, 1)
    }

    return 
}