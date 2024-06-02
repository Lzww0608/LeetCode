var mod int = 1e9 + 7
func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x 
}
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
    sum, n := 0, len(nums1)
    for i := 0; i < n; i++ {
        sum += abs(nums1[i] - nums2[i])
        //sum %= mod
    }
    arr := append([]int{}, nums1[:]...)
    sort.Ints(arr)
    res := sum
    for i := 0; i < n; i++ {
        ans := sum
        idx := sort.Search(len(arr), func(j int) bool {
            return arr[j] >= nums2[i]
        })
        p := abs(nums1[i] - nums2[i])
        if p == 0 {
            continue
        }
        t := p
        if idx < n {
            t = min(t, abs(arr[idx] - nums2[i]))
        } 
        if idx > 0 {
            t = min(t, abs(arr[idx-1] - nums2[i]))
        }
        ans = min(ans, ans - p + t)
        res = min(res, ans)
    }
    return res % mod
}
