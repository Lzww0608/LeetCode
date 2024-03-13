func maxArrayValue(nums []int) (ans int64) {

    for i := len(nums) - 1; i >= 0; i-- {
        if int64(nums[i]) <= ans {
            ans += int64(nums[i])
        } else {
            ans = int64(nums[i])
        }
    }

    return
}