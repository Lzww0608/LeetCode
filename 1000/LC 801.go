func minSwap(nums1 []int, nums2 []int) int {
    n := len(nums1)
    f := make([][2]int, n)
    for i := range f {
        f[i][0], f[i][1] = 0x3f3f3f3f, 0x3f3f3f3f
    }
    f[0][0], f[0][1] = 0, 1

    for i := 1; i < n; i++ {
        if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
            f[i][0] = f[i-1][0]
            f[i][1] = f[i-1][1] + 1
        }

        if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
            f[i][0] = min(f[i][0], f[i-1][1])
            f[i][1] = min(f[i][1], f[i-1][0] + 1)
        }
    }

    return min(f[n-1][0], f[n-1][1])
}