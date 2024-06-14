func maxScore(nums []int, k int) int64 {
    f := [2]int{}

    for i := len(nums) - 1; i >= 0; i-- {
        x := nums[i]
        r := x & 1
        f[r] = max(f[r], f[r^1] - k) + x
    }


    return int64(f[nums[0] & 1])
}