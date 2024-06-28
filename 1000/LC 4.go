func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }
    n, m := len(nums1), len(nums2)
    l, r := 0, n 
    for l <= r {
        x := l + ((r - l) >> 1)
        y := (n + m + 1) / 2 - x
        maxLeftX, minRightX := math.MinInt32, math.MaxInt32
        if x != 0 {
            maxLeftX = nums1[x-1]
        }
        if x != n {
            minRightX = nums1[x]
        }
        maxLeftY, minRightY := math.MinInt32, math.MaxInt32
        if y != 0 {
            maxLeftY = nums2[y-1]
        }
        if y != m {
            minRightY = nums2[y]
        }
        if maxLeftX <= minRightY && maxLeftY <= minRightX {
            if ((m + n) & 1) == 1 {
                return float64(max(maxLeftX, maxLeftY))
            } else {
                return float64(max(maxLeftX, maxLeftY) + min(minRightY, minRightX)) / 2.0
            }
        } else if maxLeftX > minRightY {
            r = x - 1
        } else {
            l = x + 1
        }
    }
    return 0
}
