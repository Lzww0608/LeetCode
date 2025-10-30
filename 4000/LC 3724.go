func minOperations(nums1 []int, nums2 []int) int64 {
    ans := 0
    n := len(nums1)
    for i := range n {
        ans += abs(nums1[i] - nums2[i])
    }

    x := nums2[n]
    mn := math.MaxInt32
    for i := range n {
        if nums1[i] <= x && x <= nums2[i] || nums2[i] <= x && x <= nums1[i] {
            mn = 0 
            break
        }
        mn = min(mn, min(abs(x - nums1[i]), abs(x - nums2[i])))
    }

    return int64(ans + mn + 1)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}