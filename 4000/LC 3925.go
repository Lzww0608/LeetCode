func concatWithReverse(nums []int) []int {
    n := len(nums)
    ans := make([]int, n, n * 2)
    copy(ans, nums)
    slices.Reverse(nums)
    ans = append(ans, nums...)

    return ans
}