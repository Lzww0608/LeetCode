func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)
    pre := make([]int, n + 1)
    ans := math.MaxInt32
    for i, x := range nums {
        pre[i + 1] = pre[i] + x
        for j := l; j <= r && i - j + 1 >= 0; j++ {
            if pre[i + 1] - pre[i - j + 1] > 0 {
                ans = min(ans, pre[i + 1] - pre[i - j + 1])
            }
        }
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}