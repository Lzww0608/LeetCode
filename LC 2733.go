func findNonMinOrMax(nums []int) int {
    mn, mx := nums[0], nums[0]
    for _, x := range nums {
        mn = min(mn, x)
        mx = max(mx, x)
    }

    for _, x := range nums {
        if x > mn && x < mx {
            return x
        }
    }

    return -1
}