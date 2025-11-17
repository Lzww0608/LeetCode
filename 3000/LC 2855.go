func minimumRightShifts(nums []int) int {
    swap, ans := 0, 0

    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] < nums[i - 1] {
            swap++
            ans = n - i
            if swap > 1 {
                break
            }
        }
    }

    if swap == 0 || swap == 1 && nums[0] > nums[n - 1] {
        return ans
    }
    return -1
}