func longestSubarray(nums []int, k int) (ans int) {
    cnt := make(map[int]int)

    n := len(nums)
    sum := 0
    for l, r := 0, 0; r < n; r++ {
        x := nums[r]
        if cnt[x]++; cnt[x] == 2 {
            sum++
        }

        for sum > k {
            x = nums[l]
            if cnt[x]--; cnt[x] == 1 {
                sum--
            }
            l++
        }

        ans = max(ans, r - l + 1)
    }

    return 
}