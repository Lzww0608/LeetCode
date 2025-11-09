func leftRightDifference(nums []int) []int {
    n := len(nums)
    suf := make([]int, n)
    for i := n - 2; i >= 0; i-- {
        suf[i] = suf[i + 1] + nums[i + 1]
    }

    pre := 0
    ans := make([]int, n)
    for i, x := range nums {
        ans[i] = abs(pre - suf[i])
        pre += x
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}