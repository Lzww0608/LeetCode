func minArrayLength(nums []int, k int) (ans int) {
    pre := nums[0]
    if pre == 0 {
        return 1
    }
    ans = 1
    for _, x := range nums[1:] {
        if x == 0 {
            return 1
        }
        if x * pre > k {
            ans++
            pre = x
        } else {
            pre *= x
        }
    }

    return 
}