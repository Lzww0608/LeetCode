func maximumsSplicedArray(nums1 []int, nums2 []int) int {

    var solve func([]int, []int) int
    solve = func(a, b []int) int {
        s, res, s1 := 0, 0, 0
        for i := range a {
            s1 += a[i]
            s = max(0, s + b[i] - a[i])
            res = max(res, s)
        }
        return s1 + res
    }

    return max(solve(nums1, nums2), solve(nums2, nums1))
}
