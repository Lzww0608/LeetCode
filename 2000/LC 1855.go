func maxDistance(nums1 []int, nums2 []int) (ans int) {
    n := len(nums2)
    for i, x := range nums1 {
        for i + ans < n && nums2[i + ans] >= x {
            ans++
        }
    }
    if ans > 0 {
        ans--
    }

    return
}
