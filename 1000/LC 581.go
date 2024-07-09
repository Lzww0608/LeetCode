func findUnsortedSubarray(nums []int) int {
    mx, mn := math.MinInt, math.MaxInt
    l, r := -1, -1
    n := len(nums)
    for i, x := range nums {
        if x < mx {
            r = i
        } else {
            mx = x
        }

        if nums[n-i-1] > mn {
            l = n - i - 1
        } else {
            mn = nums[n-1-i]
        }
    } 

    if r > l {
        return r - l + 1
    }

    return 0
}