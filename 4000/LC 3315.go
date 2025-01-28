func minBitwiseArray(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    for i, x := range nums {
        if x == 2 {
            ans[i] = -1
        } else {
            j := 0
            for (x >> j) & 1 == 1 {
                j++
            }
            k := (1 << j) - 1
            ans[i] = x &^ k + k >> 1
        }
    }

    return ans
}