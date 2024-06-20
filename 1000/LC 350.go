func intersect(nums1 []int, nums2 []int) (ans []int) {
    m := map[int]int{}
    for _, x := range nums1 {
        m[x]++
    }

    for _, x := range nums2 {
        if m[x] > 0 {
            ans = append(ans, x)
            m[x]--
        }
    }

    return 
}
