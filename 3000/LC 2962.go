func countSubarrays(nums []int, k int) int64 {
    mx := slices.Max(nums)
    ans, cnt := 0, 0 

    for l, r := 0, 0; r < len(nums); r++ {
        if nums[r] == mx {
            cnt++
        }

        for cnt >= k {
            if nums[l] == mx {
                cnt--
            }
            l++
        }

        ans += l
    }

    return int64(ans)
}