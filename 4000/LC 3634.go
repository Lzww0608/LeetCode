func minRemoval(nums []int, k int) int {
    n := len(nums)
    sort.Ints(nums)

    mx, l := 0, 0
    for i, x := range nums {
        for nums[l] * k < x {
            l++
        }
        mx = max(mx, i - l + 1)
    }

    return n - mx
}