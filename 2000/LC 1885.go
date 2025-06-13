func countPairs(nums1 []int, nums2 []int) (ans int64) {
    // i < j
    // (nums1[i] - nums2[i]) + (nums1[j] - nums2[j]) > 0
    n := len(nums1)
    a := make([]int, n)
    for i := range n {
        a[i] = nums1[i] - nums2[i]
    }

    sort.Ints(a)
    l, r := 0, n - 1
    for l < r {
        for l < r && a[l] + a[r] <= 0 {
            l++
        }
        ans += int64(r - l)
        r--
    }
    
    return
}