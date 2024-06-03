func minimalKSum(nums []int, k int) (ans int64) {
    nums = append(nums, 0, math.MaxInt32)
    sort.Ints(nums)
    //n := len(nums)
    for i := 1; k > 0; i++ {
        x := nums[i] - nums[i-1] - 1
        if x >= k {
            ans += int64(k * (nums[i-1] + 1 + nums[i-1] + k) / 2)
            break
        } else if x < k && x > 0 {
            ans += int64(x * (nums[i] + nums[i-1]) / 2)
            k -= x
        }
    }

    return
}