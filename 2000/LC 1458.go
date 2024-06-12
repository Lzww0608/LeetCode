func maxDotProduct(nums1 []int, nums2 []int) int {
    const MIN int = -5000000

    m, n := len(nums1), len(nums2)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        for j := range f[i] {
            f[i][j] = MIN
        }
    } 

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            prod := nums1[i] * nums2[j]
            f[i+1][j+1] = max(prod, f[i][j] + prod, f[i][j+1], f[i+1][j])
        }
    }

    return f[m][n]
}
