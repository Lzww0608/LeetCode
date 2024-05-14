func numSubarrayBoundedMax(nums []int, left int, right int) (ans int) {
    pre, cur := -1, -1
    for i, x := range nums {
        if x > right {
            pre = i
        }
        if x >= left {
            cur = i
        }
        ans += cur - pre
    }

    return
}