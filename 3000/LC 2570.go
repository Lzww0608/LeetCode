func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    m, n := len(nums1), len(nums2)
    ans := make([][]int, 0, m + n)
    i, j := 0, 0
    for i < m || j < n {
        if i >= m || j < n && nums2[j][0] < nums1[i][0] {
            ans = append(ans, nums2[j])
            j++
        } else if j >= n || i < m && nums2[j][0] > nums1[i][0] {
            ans = append(ans, nums1[i])
            i++
        } else {
            ans = append(ans, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
            i++
            j++
        }
    }

    return ans
}