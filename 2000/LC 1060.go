func missingElement(nums []int, k int) int {
    pre := nums[0] - 1
    for _, x := range nums {
        if pre + 1 < x {
            if k > x - pre - 1 {
                k -= x - pre - 1
            } else {
                return pre + k
            } 
        } 
        pre = x
    }

    return nums[len(nums)-1] + k
}