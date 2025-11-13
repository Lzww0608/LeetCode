func longestSubarray(nums []int) int {
    pre, f0, f1 := 0, 1, 1 

    ans := 1
    for i := 1; i < len(nums); i++ {
        tmp := f0
        if nums[i] >= nums[i - 1] {
            f0++
            f1++
        } else {
            f0, f1 = 1, 0
        }

        if i >= 2 && nums[i - 2] <= nums[i] {
            f1 = max(f1, pre + 2)
        } else {
            f1 = max(f1, 2)
        }

        ans = max(ans, tmp + 1, f1)
        pre = tmp
    }

    return ans
}