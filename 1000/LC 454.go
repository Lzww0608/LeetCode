func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (ans int) {
    cnt := make(map[int]int, len(nums1) * len(nums2))
    for _, x := range nums1 {
        for _, y := range nums2 {
            cnt[x + y]++
        }
    }

    for _, x := range nums3 {
        for _, y := range nums4 {
            ans += cnt[- x - y]
        }
    }

    return ans
}