func minOperations(nums1 []int, nums2 []int) int {
    ans, n, m := 0, len(nums1), len(nums2)
    if 6 * n < m || 6 * m < n {
        return -1
    }
    sum1, sum2 := 0, 0
    cnt1, cnt2 := make([]int, 6), make([]int, 6)
    for _, x := range nums1 {
        sum1 += x
        cnt1[x-1]++
    }
    for _,x := range nums2 {
        sum2 += x
        cnt2[x-1]++
    }
    if sum1 < sum2 {
        sum1, sum2 = sum2, sum1
        cnt1, cnt2 = cnt2, cnt1
    }
    d := sum1 - sum2
    for i := 0; i <= 5 && d > 0; i++ {
        for d > 0 && cnt1[5-i] > 0 {
            d -= 5 - i
            ans++
            cnt1[5-i]--
        } 
        for d > 0 && cnt2[i] > 0 {
            d -= 5 - i
            ans++
            cnt2[i]--
        }
    }
    return ans
}