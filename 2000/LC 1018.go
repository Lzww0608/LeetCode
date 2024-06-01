func prefixesDivBy5(nums []int) []bool {
    n := len(nums)
    ans := make([]bool, n)
    sum := 0
    for i, x := range nums {
        sum = (sum << 1) + x
        sum %= 5
        if sum == 0 {
            ans[i] = true
        }
    }

    return ans
}
