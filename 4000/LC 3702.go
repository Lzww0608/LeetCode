func longestSubsequence(nums []int) int {
    xor := 0
    f := false
    for _, x := range nums {
        xor ^= x
        if x != 0 {
            f = true
        }
    }

    if xor != 0 {
        return len(nums)
    } else if f {
        return len(nums) - 1
    }
    return 0
}