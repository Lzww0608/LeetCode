func maxAlternatingSum(nums []int) int64 {
    ans := 0
    for i := range nums {
        nums[i] *= nums[i]
    }
    sort.Ints(nums)
    n := len(nums)

    for i := range n {
        if i < n / 2 {
            ans -= nums[i]
        } else {
            ans += nums[i]
        }
    }

    return int64(ans)
}