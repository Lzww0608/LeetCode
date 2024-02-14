func advantageCount(nums1 []int, nums2 []int) []int {
    n := len(nums1)
    ans := make([]int, n)
    idx := make([]int, n)
    sort.Ints(nums1)
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool {
        return nums2[idx[i]] < nums2[idx[j]]
    })
    i, j := 0, n - 1
    for _, x := range nums1 {
        if x > nums2[idx[i]] {
            ans[idx[i]] = x
            i++
        } else {
            ans[idx[j]] = x
            j--
        }
    }
    return ans
}