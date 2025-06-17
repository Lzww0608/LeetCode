func numberOfArithmeticSlices(nums []int) (ans int) {
    n := len(nums)
    for l, r := 0, 2; r < n; r++ {
        if nums[r] - nums[r-1] == nums[r-1] - nums[r-2] {
            ans += r - l - 1
        } else {
            l = r - 1
        }
    }

    return
}