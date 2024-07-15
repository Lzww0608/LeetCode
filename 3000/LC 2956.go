func findIntersectionValues(nums1 []int, nums2 []int) []int {
    a, b := 0, 0
    m1 := map[int]bool{}
    m2 := map[int]bool{}

    for _, x := range nums1 {
        m1[x] = true
    }

    for _, x := range nums2 {
        if m1[x] {
            b++
        }
        m2[x] = true
    }

    for _, x := range nums1 {
        if m2[x] {
            a++
        }
    }

    return []int{a, b}
}