func longestOnes(nums []int, k int) (ans int) {
    n := len(nums)
    cnt := 0
    for l, r := 0, 0; r < n; r++ {
        if nums[r] == 0 {
            cnt++
        }
        for cnt > k {
            if nums[l] == 0 {
                cnt--
            }
            l++
        }

        ans = max(ans, r - l + 1)
    }

    return
}