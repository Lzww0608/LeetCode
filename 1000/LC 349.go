func intersection(nums1 []int, nums2 []int) (ans []int) {
    m := map[int]bool{}
    for _, x := range nums1 {
        m[x] = true
    }

    for _, x := range nums2 {
        if m[x] {
            m[x] = false
            ans = append(ans, x)
        }
    }

    return 
}
